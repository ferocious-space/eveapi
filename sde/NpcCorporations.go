// Code generated by YAML2GO. DO NOT EDIT.

package sde

import (
	yamlv3 "gopkg.in/yaml.v3"
	"os"
)

type NpcCorporationMap map[int32]NpcCorporation

func (x *NpcCorporationMap) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yamlv3.NewDecoder(f).Decode(x)
}
func (x NpcCorporationMap) Get(ID int32) *NpcCorporation {
	if a, ok := x[ID]; ok {
		return &a
	}
	return nil
}

type NpcCorporation struct {
	CeoID                      *int32                            `bson:"ceoID,omitempty" db:"ceoID,omitempty" json:"ceoID,omitempty" yaml:"ceoID,omitempty"`
	Deleted                    *bool                             `bson:"deleted,omitempty" db:"deleted,omitempty" json:"deleted,omitempty" yaml:"deleted,omitempty"`
	DescriptionID              *NpcCorporationDescriptionID      `bson:"descriptionID,omitempty" db:"descriptionID,omitempty" json:"descriptionID,omitempty" yaml:"descriptionID,omitempty"`
	Extent                     *string                           `bson:"extent,omitempty" db:"extent,omitempty" json:"extent,omitempty" yaml:"extent,omitempty"`
	HasPlayerPersonnelManager  *bool                             `bson:"hasPlayerPersonnelManager,omitempty" db:"hasPlayerPersonnelManager,omitempty" json:"hasPlayerPersonnelManager,omitempty" yaml:"hasPlayerPersonnelManager,omitempty"`
	InitialPrice               *int32                            `bson:"initialPrice,omitempty" db:"initialPrice,omitempty" json:"initialPrice,omitempty" yaml:"initialPrice,omitempty"`
	MemberLimit                *int32                            `bson:"memberLimit,omitempty" db:"memberLimit,omitempty" json:"memberLimit,omitempty" yaml:"memberLimit,omitempty"`
	MinSecurity                *float64                          `bson:"minSecurity,omitempty" db:"minSecurity,omitempty" json:"minSecurity,omitempty" yaml:"minSecurity,omitempty"`
	MinimumJoinStanding        *int32                            `bson:"minimumJoinStanding,omitempty" db:"minimumJoinStanding,omitempty" json:"minimumJoinStanding,omitempty" yaml:"minimumJoinStanding,omitempty"`
	NameID                     *NpcCorporationNameID             `bson:"nameID,omitempty" db:"nameID,omitempty" json:"nameID,omitempty" yaml:"nameID,omitempty"`
	PublicShares               *int32                            `bson:"publicShares,omitempty" db:"publicShares,omitempty" json:"publicShares,omitempty" yaml:"publicShares,omitempty"`
	SendCharTerminationMessage *bool                             `bson:"sendCharTerminationMessage,omitempty" db:"sendCharTerminationMessage,omitempty" json:"sendCharTerminationMessage,omitempty" yaml:"sendCharTerminationMessage,omitempty"`
	Shares                     *int32                            `bson:"shares,omitempty" db:"shares,omitempty" json:"shares,omitempty" yaml:"shares,omitempty"`
	Size                       *string                           `bson:"size,omitempty" db:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
	StationID                  *int32                            `bson:"stationID,omitempty" db:"stationID,omitempty" json:"stationID,omitempty" yaml:"stationID,omitempty"`
	TaxRate                    *float64                          `bson:"taxRate,omitempty" db:"taxRate,omitempty" json:"taxRate,omitempty" yaml:"taxRate,omitempty"`
	TickerName                 *string                           `bson:"tickerName,omitempty" db:"tickerName,omitempty" json:"tickerName,omitempty" yaml:"tickerName,omitempty"`
	UniqueName                 *bool                             `bson:"uniqueName,omitempty" db:"uniqueName,omitempty" json:"uniqueName,omitempty" yaml:"uniqueName,omitempty"`
	AllowedMemberRaces         []int32                           `bson:"allowedMemberRaces,omitempty" db:"allowedMemberRaces,omitempty" json:"allowedMemberRaces,omitempty" yaml:"allowedMemberRaces,omitempty"`
	CorporationTrades          map[int32]float64                 `bson:"corporationTrades,omitempty" db:"corporationTrades,omitempty" json:"corporationTrades,omitempty" yaml:"corporationTrades,omitempty"`
	Divisions                  map[int32]NpcCorporationDivisions `bson:"divisions,omitempty" db:"divisions,omitempty" json:"divisions,omitempty" yaml:"divisions,omitempty"`
	EnemyID                    *int32                            `bson:"enemyID,omitempty" db:"enemyID,omitempty" json:"enemyID,omitempty" yaml:"enemyID,omitempty"`
	FactionID                  *int32                            `bson:"factionID,omitempty" db:"factionID,omitempty" json:"factionID,omitempty" yaml:"factionID,omitempty"`
	FriendID                   *int32                            `bson:"friendID,omitempty" db:"friendID,omitempty" json:"friendID,omitempty" yaml:"friendID,omitempty"`
	IconID                     *int32                            `bson:"iconID,omitempty" db:"iconID,omitempty" json:"iconID,omitempty" yaml:"iconID,omitempty"`
	Investors                  map[int32]int32                   `bson:"investors,omitempty" db:"investors,omitempty" json:"investors,omitempty" yaml:"investors,omitempty"`
	LpOfferTables              []int32                           `bson:"lpOfferTables,omitempty" db:"lpOfferTables,omitempty" json:"lpOfferTables,omitempty" yaml:"lpOfferTables,omitempty"`
	MainActivityID             *int32                            `bson:"mainActivityID,omitempty" db:"mainActivityID,omitempty" json:"mainActivityID,omitempty" yaml:"mainActivityID,omitempty"`
	RaceID                     *int32                            `bson:"raceID,omitempty" db:"raceID,omitempty" json:"raceID,omitempty" yaml:"raceID,omitempty"`
	SizeFactor                 *float64                          `bson:"sizeFactor,omitempty" db:"sizeFactor,omitempty" json:"sizeFactor,omitempty" yaml:"sizeFactor,omitempty"`
	SolarSystemID              *int32                            `bson:"solarSystemID,omitempty" db:"solarSystemID,omitempty" json:"solarSystemID,omitempty" yaml:"solarSystemID,omitempty"`
	SecondaryActivityID        *int32                            `bson:"secondaryActivityID,omitempty" db:"secondaryActivityID,omitempty" json:"secondaryActivityID,omitempty" yaml:"secondaryActivityID,omitempty"`
	Url                        *string                           `bson:"url,omitempty" db:"url,omitempty" json:"url,omitempty" yaml:"url,omitempty"`
	ExchangeRates              map[int32]float64                 `bson:"exchangeRates,omitempty" db:"exchangeRates,omitempty" json:"exchangeRates,omitempty" yaml:"exchangeRates,omitempty"`
}
type NpcCorporationDescriptionID struct {
	De *string `bson:"de,omitempty" db:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" db:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Es *string `bson:"es,omitempty" db:"es,omitempty" json:"es,omitempty" yaml:"es,omitempty"`
	Fr *string `bson:"fr,omitempty" db:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	It *string `bson:"it,omitempty" db:"it,omitempty" json:"it,omitempty" yaml:"it,omitempty"`
	Ja *string `bson:"ja,omitempty" db:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" db:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" db:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" db:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
type NpcCorporationDivisions struct {
	DivisionNumber *int32 `bson:"divisionNumber,omitempty" db:"divisionNumber,omitempty" json:"divisionNumber,omitempty" yaml:"divisionNumber,omitempty"`
	LeaderID       *int32 `bson:"leaderID,omitempty" db:"leaderID,omitempty" json:"leaderID,omitempty" yaml:"leaderID,omitempty"`
	Size           *int32 `bson:"size,omitempty" db:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty"`
}
type NpcCorporationNameID struct {
	De *string `bson:"de,omitempty" db:"de,omitempty" json:"de,omitempty" yaml:"de,omitempty"`
	En *string `bson:"en,omitempty" db:"en,omitempty" json:"en,omitempty" yaml:"en,omitempty"`
	Es *string `bson:"es,omitempty" db:"es,omitempty" json:"es,omitempty" yaml:"es,omitempty"`
	Fr *string `bson:"fr,omitempty" db:"fr,omitempty" json:"fr,omitempty" yaml:"fr,omitempty"`
	It *string `bson:"it,omitempty" db:"it,omitempty" json:"it,omitempty" yaml:"it,omitempty"`
	Ja *string `bson:"ja,omitempty" db:"ja,omitempty" json:"ja,omitempty" yaml:"ja,omitempty"`
	Ko *string `bson:"ko,omitempty" db:"ko,omitempty" json:"ko,omitempty" yaml:"ko,omitempty"`
	Ru *string `bson:"ru,omitempty" db:"ru,omitempty" json:"ru,omitempty" yaml:"ru,omitempty"`
	Zh *string `bson:"zh,omitempty" db:"zh,omitempty" json:"zh,omitempty" yaml:"zh,omitempty"`
}
