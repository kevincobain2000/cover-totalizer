package pkg

import (
	"encoding/xml"
	"fmt"
	"math"
	"os"
	"strconv"
)

type CoverageXMLJacoco struct {
	Counter []struct {
		Text    string `xml:",chardata"`
		Type    string `xml:"type,attr"`
		Missed  string `xml:"missed,attr"`
		Covered string `xml:"covered,attr"`
	} `xml:"counter"`
}
type CoverageXMLGO struct {
	LineRate float64 `xml:"line-rate,attr"`
}
type CoverageXMLPHP struct {
	Project struct {
		Metrics struct {
			Statements        int64 `xml:"statements,attr"`
			Coveredstatements int64 `xml:"coveredstatements,attr"`
		} `xml:"metrics"`
	} `xml:"project"`
}

// CoverageService is the service for Uploading coverage assets from user
type CoverageService struct {
}

// NewCoverageService creates a new CoverageService
func NewCoverageService() *CoverageService {
	return &CoverageService{}
}

func (s *CoverageService) ParseCoveragePercentage(coverageXMLPath string) (float64, error) {
	// try with Clover. Even if the format is not clover, it will be parsed without errors.
	percentage, err := s.parseClover(coverageXMLPath)
	if err != nil {
		return 0.0, err
	}
	if percentage > 0.0 {
		return percentage, nil
	}

	// Since previous parse was not able to parse the coverage, try with PHPUnit.
	percentage, err = s.parsePHPUnit(coverageXMLPath)
	if err != nil {
		return 0.0, err
	}

	if percentage > 0.0 {
		return percentage, nil
	}

	// Since previous parse was not able to parse the coverage, try with PHPUnit.
	percentage, err = s.parseJacoco(coverageXMLPath)
	if err != nil {
		return 0.0, err
	}
	if percentage > 0.0 {
		return percentage, nil
	}

	err = fmt.Errorf("unable to parse coverage file")
	return 0.0, err
}

func (s *CoverageService) parseClover(coverageXMLPath string) (float64, error) {
	//#nosec G304
	data, err := os.ReadFile(coverageXMLPath)
	if err != nil {
		return 0.0, err
	}
	coverageXMLGO := &CoverageXMLGO{}
	err = xml.Unmarshal([]byte(data), &coverageXMLGO)
	if err != nil {
		return 0.0, err
	}

	percentage := coverageXMLGO.LineRate * 100
	return roundNearest(percentage), nil
}

func (s *CoverageService) parseJacoco(coverageXMLPath string) (float64, error) {
	//#nosec G304
	data, err := os.ReadFile(coverageXMLPath)
	if err != nil {
		return 0.0, err
	}
	coverageXMLJacoco := &CoverageXMLJacoco{}
	err = xml.Unmarshal([]byte(data), &coverageXMLJacoco)
	if err != nil {
		return 0.0, err
	}
	for _, counter := range coverageXMLJacoco.Counter {
		if counter.Type == "LINE" {
			covered, _ := strconv.ParseInt(counter.Covered, 10, 64)
			missed, _ := strconv.ParseInt(counter.Missed, 10, 64)
			percentage := float64(covered) / float64(covered+missed) * 100
			return roundNearest(percentage), nil
		}
	}

	return 0.0, nil
}

func (s *CoverageService) parsePHPUnit(coverageXMLPath string) (float64, error) {
	//#nosec G304
	data, err := os.ReadFile(coverageXMLPath)
	if err != nil {
		return 0.0, err
	}
	coverageXMLPHP := &CoverageXMLPHP{}
	err = xml.Unmarshal([]byte(data), &coverageXMLPHP)
	if err != nil {
		return 0.0, err
	}

	if coverageXMLPHP.Project.Metrics.Statements == 0 {
		return 0.0, nil
	}

	percentage := float64(coverageXMLPHP.Project.Metrics.Coveredstatements) / float64(coverageXMLPHP.Project.Metrics.Statements) * 100
	return roundNearest(percentage), nil
}

func roundNearest(x float64) float64 {
	return math.Round(x*100) / 100
}
