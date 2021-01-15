package model

import (
	"fmt"
	"reflect"
	"strconv"
)

type (
	//PlayerBasicBox represents a players basic statlines in a game
	PlayerBasicBox struct {
		Name      string
		MP        string  `bref:"mp"`
		FG        int     `bref:"fg"`
		FGA       int     `bref:"fga"`
		FGPct     float64 `bref:"fg_pct"`
		ThreeP    int     `bref:"fg3"`
		ThreePA   int     `bref:"fg3a"`
		ThreePPct float64 `bref:"fg3_pct"`
		FT        int     `bref:"ft"`
		FTA       int     `bref:"fta"`
		FTPct     float64 `bref:"ft_pct"`
		ORB       int     `bref:"orb"`
		DRB       int     `bref:"drb"`
		TRB       int     `bref:"trb"`
		AST       int     `bref:"ast"`
		STL       int     `bref:"stl"`
		BLK       int     `bref:"blk"`
		TOV       int     `bref:"tov"`
		PF        int     `bref:"pf"`
		PTS       int     `bref:"pts"`
		PlusMin   string  `bref:"plus_minus"`
		DNP       string  `bref:"reason"`
		SourceID  string
	}

	//PlayerAdvBox represents a players advanced stats during a game.
	PlayerAdvBox struct {
		Name         string
		MP           string  `bref:"mp"`
		TrueShootPct float64 `bref:"ts_pct"`
		EFG          float64 `bref:"efg_pct"`
		ThreePARate  float64 `bref:"fg3a_per_fga_pct"`
		FTARate      float64 `bref:"fta_per_fga_pct"`
		ORBPct       float64 `bref:"orb_pct"`
		DRBPct       float64 `bref:"drb_pct"`
		TRBPct       float64 `bref:"trb_pct"`
		ASTPct       float64 `bref:"ast_pct"`
		STLPct       float64 `bref:"stl_pct"`
		BLKPct       float64 `bref:"blk_pct"`
		TOVPct       float64 `bref:"tov_pct"`
		USGRate      float64 `bref:"usg_pct"`
		ORTG         int     `bref:"off_rtg"`
		DRTG         int     `bref:"def_rtg"`
		BPM          string  `bref:"bpm"`
		DNP          string  `bref:"reason"`
		SourceID     string
	}
)

//AddByTag will search the struct's fields for a field with the tag <tag>
// and will set the value of that field to <value>. An error will be returned if
// a filed with the given tag does not exist.
func (p *PlayerAdvBox) AddByTag(tag, value string) error {

	// this is probably not neccesary, but I was messing around with different options
	// and dont want to rewrite something works right now.
	reflectRead := reflect.ValueOf(*p)
	reflectWrite := reflect.ValueOf(p)
	//iterate over struct to find Field with a matching tag.
	found := false
Find:
	for i := 0; i < reflectRead.NumField(); i++ {
		field := reflectWrite.Elem().Field(i)
		fieldTag, _ := reflectRead.Type().Field(i).Tag.Lookup("bref")
		if fieldTag == tag {
			switch field.Kind() {
			case reflect.Int:
				typeVal, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				field.SetInt(typeVal)
			case reflect.Float64:
				typeVal, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				field.SetFloat(typeVal)
			default:
				field.SetString(value)
			}
			found = true
			break Find
		}

	}

	if !found {
		return fmt.Errorf("could dont find field with tag %v", tag)
	}

	return nil
}

//AddByTag will search the struct's fields for a field with the tag <tag>
// and will set the value of that field to <value>. An error will be returned if
// a filed with the given tag does not exist.
func (p *PlayerBasicBox) AddByTag(tag, value string) error {

	// this is probably not neccesary, but I was messing around with different options
	// and dont want to rewrite something works right now.
	reflectRead := reflect.ValueOf(*p)
	reflectWrite := reflect.ValueOf(p)
	//iterate over struct to find Field with a matching tag.
	found := false
Find:
	for i := 0; i < reflectRead.NumField(); i++ {
		field := reflectWrite.Elem().Field(i)
		fieldTag, _ := reflectRead.Type().Field(i).Tag.Lookup("bref")
		if fieldTag == tag {
			//convert string value to its proper type.
			switch field.Kind() {
			case reflect.Int:
				typeVal, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				field.SetInt(typeVal)
			case reflect.Float64:
				typeVal, err := strconv.ParseFloat(value, 64)
				if err != nil {
					return err
				}
				field.SetFloat(typeVal)
			default:
				field.SetString(value)
			}
			found = true
			break Find
		}

	}

	if !found {
		return fmt.Errorf("could dont find field with tag %v", tag)
	}

	return nil
}
