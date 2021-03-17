/*
 *    Copyright 2021 FerociousBite and Contributors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package eveapi

import (
	"net/http"
	"net/http/cookiejar"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"go.uber.org/zap"
	"golang.org/x/net/publicsuffix"

	"github.com/ferocious-space/durableclient"
	"github.com/ferocious-space/httpcache"

	"github.com/ferocious-space/eveapi/esi"
)

func NewAPIClient(cache httpcache.Cache, logger *zap.Logger) *esi.EVESwaggerInterface {

	esiHTTPClient := durableclient.NewClient("ESI", "https://github.com/ferocious-space/eveapi", logger.Named("ESI"))
	jar, err := cookiejar.New(&cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	})
	if err != nil {
		return nil
	}
	var cacheTransport *httpcache.Transport

	// ensure there is at least SOME caching
	if cache != nil {
		cacheTransport = httpcache.NewTransport(cache)
	} else {
		cacheTransport = httpcache.NewTransport(httpcache.NewLRUCache(1<<20*128, 300))
	}

	// setting our custom HTTP Client AFTER the cache , only cache misses will go to HTTP client. we dont want to retry on cache
	cacheTransport.Transport = esiHTTPClient.Transport

	apiRuntime := httptransport.NewWithClient(esi.DefaultHost, esi.DefaultBasePath, esi.DefaultSchemes, &http.Client{
		Transport: cacheTransport,
		Jar:       jar,
	})

	return esi.New(apiRuntime, strfmt.Default)
}
