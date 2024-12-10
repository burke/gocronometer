package gocronometer

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type ServingRecord struct {
	RecordedTime     time.Time
	Group            string
	FoodName         string
	QuantityValue    float64
	QuantityUnits    string
	EnergyKcal       float64
	CaffeineMg       float64
	WaterG           float64
	B1Mg             float64
	B2Mg             float64
	B3Mg             float64
	B5Mg             float64
	B6Mg             float64
	B12Mg            float64
	BiotinUg         float64
	CholineMg        float64
	FolateUg         float64
	VitaminAUI       float64
	VitaminCMg       float64
	VitaminDUI       float64
	VitaminEMg       float64
	VitaminKMg       float64
	CalciumMg        float64
	ChromiumUg       float64
	CopperMg         float64
	FluorideUg       float64
	IodineUg         float64
	MagnesiumMg      float64
	ManganeseMg      float64
	PhosphorusMg     float64
	PotassiumMg      float64
	SeleniumUg       float64
	SodiumMg         float64
	ZincMg           float64
	CarbsG           float64
	FiberG           float64
	FructoseG        float64
	GalactoseG       float64
	GlucoseG         float64
	LactoseG         float64
	MaltoseG         float64
	StarchG          float64
	SucroseG         float64
	SugarsG          float64
	NetCarbsG        float64
	FatG             float64
	CholesterolMg    float64
	MonounsaturatedG float64
	PolyunsaturatedG float64
	SaturatedG       float64
	TransFatG        float64
	Omega3G          float64
	Omega6G          float64
	CystineG         float64
	HistidineG       float64
	IsoleucineG      float64
	LeucineG         float64
	LysineG          float64
	MethionineG      float64
	PhenylalanineG   float64
	ThreonineG       float64
	TryptophanG      float64
	TyrosineG        float64
	ValineG          float64
	ProteinG         float64
	IronMg           float64
	Category         string
}

type ServingRecords []ServingRecord

type ServingsExport struct {
	Records ServingRecords
}

const (
	DateTimeFormat = "2006-01-02 15:04"
)

// parseDateTime handles parsing of Cronometer date+time strings
func parseDateTime(date, timeStr string, location *time.Location) (time.Time, error) {
	if location == nil {
		location = time.UTC
	}

	// Default to midnight if no time provided
	if timeStr == "" {
		timeStr = "00:00"
	}

	// Combine date and time
	dateTimeStr := date + " " + timeStr

	// Parse with location
	t, err := time.ParseInLocation(DateTimeFormat, dateTimeStr, location)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date/time format %q: %w", dateTimeStr, err)
	}

	return t, nil
}

func ParseServingsExport(rawCSVReader io.Reader, location *time.Location) (ServingRecords, error) {

	r := csv.NewReader(rawCSVReader)

	lineNum := 0
	headers := make(map[int]string)
	servings := make(ServingRecords, 0, 0)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Index all the headers.
		if lineNum == 0 {

			for i, v := range record {
				headers[i] = v
			}
			lineNum++
			continue
		}
		lineNum++

		var date string
		var timeStr string
		serving := ServingRecord{}
		for i, v := range record {
			columnName := headers[i]

			switch columnName {
			case "Day":
				date = v
			case "Time":
				timeStr = v
			case "Group":
				serving.Group = v
			case "Food Name":
				serving.FoodName = v
			case "Amount":
				parts := strings.SplitN(v, " ", 2)
				if len(parts) < 2 {
					return nil, fmt.Errorf("invalid amount format %q, expected 'value unit'", v)
				}
				f, err := parseFloat(parts[0], 64)
				if err != nil {
					return nil, fmt.Errorf("parsing quantity value %q: %w", parts[0], err)
				}
				serving.QuantityValue = f
				serving.QuantityUnits = parts[1]
			case "Energy (kcal)":
				f, err := parseNutrientFloat(v, "energy")
				if err != nil {
					return nil, err
				}
				serving.EnergyKcal = f
			case "Caffeine (mg)":
				f, err := parseNutrientFloat(v, "caffeine")
				if err != nil {
					return nil, err
				}
				serving.CaffeineMg = f
			case "Water (g)":
				f, err := parseNutrientFloat(v, "water")
				if err != nil {
					return nil, err
				}
				serving.WaterG = f
			case "B1 (Thiamine) (mg)":
				f, err := parseNutrientFloat(v, "vitamin B1")
				if err != nil {
					return nil, err
				}
				serving.B1Mg = f
			case "B2 (Riboflavin) (mg)":
				f, err := parseNutrientFloat(v, "vitamin B2")
				if err != nil {
					return nil, err
				}
				serving.B2Mg = f
			case "B3 (Niacin) (mg)":
				f, err := parseNutrientFloat(v, "vitamin B3")
				if err != nil {
					return nil, err
				}
				serving.B3Mg = f
			case "B5 (Pantothenic Acid) (mg)":
				f, err := parseNutrientFloat(v, "vitamin B5")
				if err != nil {
					return nil, err
				}
				serving.B5Mg = f
			case "B6 (Pyridoxine) (mg)":
				f, err := parseNutrientFloat(v, "vitamin B6")
				if err != nil {
					return nil, err
				}
				serving.B6Mg = f
			case "B12 (Cobalamin) (µg)":
				f, err := parseNutrientFloat(v, "vitamin B12")
				if err != nil {
					return nil, err
				}
				serving.B12Mg = f
			case "Biotin (µg)":
				f, err := parseNutrientFloat(v, "biotin")
				if err != nil {
					return nil, err
				}
				serving.BiotinUg = f
			case "Choline (mg)":
				f, err := parseNutrientFloat(v, "choline")
				if err != nil {
					return nil, err
				}
				serving.CholineMg = f
			case "Folate (µg)":
				f, err := parseNutrientFloat(v, "folate")
				if err != nil {
					return nil, err
				}
				serving.FolateUg = f
			case "Vitamin A (IU)":
				f, err := parseNutrientFloat(v, "vitamin A")
				if err != nil {
					return nil, err
				}
				serving.VitaminAUI = f
			case "Vitamin C (mg)":
				f, err := parseNutrientFloat(v, "vitamin C")
				if err != nil {
					return nil, err
				}
				serving.VitaminCMg = f
			case "Vitamin D (IU)":
				f, err := parseNutrientFloat(v, "vitamin D")
				if err != nil {
					return nil, err
				}
				serving.VitaminDUI = f
			case "Vitamin E (mg)":
				f, err := parseNutrientFloat(v, "vitamin E")
				if err != nil {
					return nil, err
				}
				serving.VitaminEMg = f
			case "Vitamin K (µg)":
				f, err := parseNutrientFloat(v, "vitamin K")
				if err != nil {
					return nil, err
				}
				serving.VitaminKMg = f
			case "Calcium (mg)":
				f, err := parseNutrientFloat(v, "calcium")
				if err != nil {
					return nil, err
				}
				serving.CalciumMg = f
			case "Chromium (µg)":
				f, err := parseNutrientFloat(v, "chromium")
				if err != nil {
					return nil, err
				}
				serving.ChromiumUg = f
			case "Copper (mg)":
				f, err := parseNutrientFloat(v, "copper")
				if err != nil {
					return nil, err
				}
				serving.CopperMg = f
			case "Fluoride (µg)":
				f, err := parseNutrientFloat(v, "fluoride")
				if err != nil {
					return nil, err
				}
				serving.FluorideUg = f
			case "Iodine (µg)":
				f, err := parseNutrientFloat(v, "iodine")
				if err != nil {
					return nil, err
				}
				serving.IodineUg = f
			case "Iron (mg)":
				f, err := parseNutrientFloat(v, "iron")
				if err != nil {
					return nil, err
				}
				serving.IronMg = f
			case "Magnesium (mg)":
				f, err := parseNutrientFloat(v, "magnesium")
				if err != nil {
					return nil, err
				}
				serving.MagnesiumMg = f
			case "Manganese (mg)":
				f, err := parseNutrientFloat(v, "manganese")
				if err != nil {
					return nil, err
				}
				serving.ManganeseMg = f
			case "Phosphorus (mg)":
				f, err := parseNutrientFloat(v, "phosphorus")
				if err != nil {
					return nil, err
				}
				serving.PhosphorusMg = f
			case "Potassium (mg)":
				f, err := parseNutrientFloat(v, "potassium")
				if err != nil {
					return nil, err
				}
				serving.PotassiumMg = f
			case "Selenium (µg)":
				f, err := parseNutrientFloat(v, "selenium")
				if err != nil {
					return nil, err
				}
				serving.SeleniumUg = f
			case "Sodium (mg)":
				f, err := parseNutrientFloat(v, "sodium")
				if err != nil {
					return nil, err
				}
				serving.SodiumMg = f
			case "Zinc (mg)":
				f, err := parseNutrientFloat(v, "zinc")
				if err != nil {
					return nil, err
				}
				serving.ZincMg = f
			case "Carbs (g)":
				f, err := parseNutrientFloat(v, "carbohydrates")
				if err != nil {
					return nil, err
				}
				serving.CarbsG = f
			case "Fiber (g)":
				f, err := parseNutrientFloat(v, "fiber")
				if err != nil {
					return nil, err
				}
				serving.FiberG = f
			case "Fructose (g)":
				f, err := parseNutrientFloat(v, "fructose")
				if err != nil {
					return nil, err
				}
				serving.FructoseG = f
			case "Galactose (g)":
				f, err := parseNutrientFloat(v, "galactose")
				if err != nil {
					return nil, err
				}
				serving.GalactoseG = f
			case "Glucose (g)":
				f, err := parseNutrientFloat(v, "glucose")
				if err != nil {
					return nil, err
				}
				serving.GlucoseG = f
			case "Lactose (g)":
				f, err := parseNutrientFloat(v, "lactose")
				if err != nil {
					return nil, err
				}
				serving.LactoseG = f
			case "Maltose (g)":
				f, err := parseNutrientFloat(v, "maltose")
				if err != nil {
					return nil, err
				}
				serving.MaltoseG = f
			case "Starch (g)":
				f, err := parseNutrientFloat(v, "starch")
				if err != nil {
					return nil, err
				}
				serving.StarchG = f
			case "Sucrose (g)":
				f, err := parseNutrientFloat(v, "sucrose")
				if err != nil {
					return nil, err
				}
				serving.SucroseG = f
			case "Sugars (g)":
				f, err := parseNutrientFloat(v, "sugars")
				if err != nil {
					return nil, err
				}
				serving.SugarsG = f
			case "Net Carbs (g)":
				f, err := parseNutrientFloat(v, "net carbs")
				if err != nil {
					return nil, err
				}
				serving.NetCarbsG = f
			case "Fat (g)":
				f, err := parseNutrientFloat(v, "fat")
				if err != nil {
					return nil, err
				}
				serving.FatG = f
			case "Cholesterol (mg)":
				f, err := parseNutrientFloat(v, "cholesterol")
				if err != nil {
					return nil, err
				}
				serving.CholesterolMg = f
			case "Monounsaturated (g)":
				f, err := parseNutrientFloat(v, "monounsaturated fat")
				if err != nil {
					return nil, err
				}
				serving.MonounsaturatedG = f
			case "Polyunsaturated (g)":
				f, err := parseNutrientFloat(v, "polyunsaturated fat")
				if err != nil {
					return nil, err
				}
				serving.PolyunsaturatedG = f
			case "Saturated (g)":
				f, err := parseNutrientFloat(v, "saturated fat")
				if err != nil {
					return nil, err
				}
				serving.SaturatedG = f
			case "Trans-Fats (g)":
				f, err := parseNutrientFloat(v, "trans fat")
				if err != nil {
					return nil, err
				}
				serving.TransFatG = f
			case "Omega-3 (g)":
				f, err := parseNutrientFloat(v, "omega-3")
				if err != nil {
					return nil, err
				}
				serving.Omega3G = f
			case "Omega-6 (g)":
				f, err := parseNutrientFloat(v, "omega-6")
				if err != nil {
					return nil, err
				}
				serving.Omega6G = f
			case "Cystine (g)":
				f, err := parseNutrientFloat(v, "cystine")
				if err != nil {
					return nil, err
				}
				serving.CystineG = f
			case "Histidine (g)":
				f, err := parseNutrientFloat(v, "histidine")
				if err != nil {
					return nil, err
				}
				serving.HistidineG = f
			case "Isoleucine (g)":
				f, err := parseNutrientFloat(v, "isoleucine")
				if err != nil {
					return nil, err
				}
				serving.IsoleucineG = f
			case "Leucine (g)":
				f, err := parseNutrientFloat(v, "leucine")
				if err != nil {
					return nil, err
				}
				serving.LeucineG = f
			case "Lysine (g)":
				f, err := parseNutrientFloat(v, "lysine")
				if err != nil {
					return nil, err
				}
				serving.LysineG = f
			case "Methionine (g)":
				f, err := parseNutrientFloat(v, "methionine")
				if err != nil {
					return nil, err
				}
				serving.MethionineG = f
			case "Phenylalanine (g)":
				f, err := parseNutrientFloat(v, "phenylalanine")
				if err != nil {
					return nil, err
				}
				serving.PhenylalanineG = f
			case "Protein (g)":
				f, err := parseNutrientFloat(v, "protein")
				if err != nil {
					return nil, err
				}
				serving.ProteinG = f
			case "Threonine (g)":
				f, err := parseNutrientFloat(v, "threonine")
				if err != nil {
					return nil, err
				}
				serving.ThreonineG = f
			case "Tryptophan (g)":
				f, err := parseNutrientFloat(v, "tryptophan")
				if err != nil {
					return nil, err
				}
				serving.TryptophanG = f
			case "Tyrosine (g)":
				f, err := parseNutrientFloat(v, "tyrosine")
				if err != nil {
					return nil, err
				}
				serving.TyrosineG = f
			case "Valine (g)":
				f, err := parseNutrientFloat(v, "valine")
				if err != nil {
					return nil, err
				}
				serving.ValineG = f
			case "Category":
				serving.Category = v
			}

		}
		if timeStr == "" {
			timeStr = "00:00 AM"
		}

		if location == nil {
			location = time.UTC
		}

		serving.RecordedTime, err = parseDateTime(date, timeStr, location)
		if err != nil {
			return nil, fmt.Errorf("parsing serving time: %w", err)
		}
		servings = append(servings, serving)
	}

	return servings, nil

}

// parseFloat wraps time.ParseFloat but interprites an empty string as 0.
func parseFloat(s string, bitSize int) (float64, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.ParseFloat(s, bitSize)
}

type ExerciseRecord struct {
	RecordedTime   time.Time
	Exercise       string
	Minutes        float64
	CaloriesBurned float64
}

type ExerciseRecords []ExerciseRecord

func ParseExerciseExport(rawCSVReader io.Reader, location *time.Location) (ExerciseRecords, error) {

	r := csv.NewReader(rawCSVReader)

	lineNum := 0
	headers := make(map[int]string)
	exercises := make(ExerciseRecords, 0, 0)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Index all the headers.
		if lineNum == 0 {

			for i, v := range record {
				headers[i] = v
			}
			lineNum++
			continue
		}
		lineNum++

		var date string
		var timeStr string
		exercise := ExerciseRecord{}
		for i, v := range record {
			columnName := headers[i]

			switch columnName {
			case "Day":
				date = v
			case "Time":
				timeStr = v
			case "Exercise":
				exercise.Exercise = v
			case "Minutes":
				f, err := parseFloat(v, 64)
				if err != nil {
					return nil, fmt.Errorf("parsing energy: %s", err)
				}
				exercise.Minutes = f

			case "Calories Burned":
				f, err := parseFloat(v, 64)
				if err != nil {
					return nil, fmt.Errorf("parsing caffeine: %s", err)
				}
				exercise.CaloriesBurned = f

			}
		}
		if timeStr == "" {
			timeStr = "00:00 AM"
		}

		if location == nil {
			location = time.UTC
		}
		exercise.RecordedTime, err = parseDateTime(date, timeStr, location)
		if err != nil {
			return nil, fmt.Errorf("parsing exercise time: %w", err)
		}
		exercises = append(exercises, exercise)
	}

	return exercises, nil

}

type BiometricRecord struct {
	RecordedTime time.Time
	Metric       string
	Unit         string
	Amount       float64
}

type BiometricRecords []BiometricRecord

func ParseBiometricRecordsExport(rawCSVReader io.Reader, location *time.Location) (BiometricRecords, error) {

	r := csv.NewReader(rawCSVReader)

	lineNum := 0
	headers := make(map[int]string)
	records := make(BiometricRecords, 0, 0)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Index all the headers.
		if lineNum == 0 {

			for i, v := range record {
				headers[i] = v
			}
			lineNum++
			continue
		}
		lineNum++

		var date string
		var timeStr string
		bioRecord := BiometricRecord{}
		for i, v := range record {
			columnName := headers[i]

			switch columnName {
			case "Day":
				date = v
			case "Time":
				timeStr = v
			case "Metric":
				bioRecord.Metric = v
			case "Unit":
				bioRecord.Unit = v
			case "Amount":
				if !strings.Contains(v, "/") {
					f, err := parseFloat(v, 64)
					if err != nil {
						return nil, fmt.Errorf("parsing energy: %s", err)
					}
					bioRecord.Amount = f
				}
			}
		}
		if timeStr == "" {
			timeStr = "00:00 AM"
		}

		if location == nil {
			location = time.UTC
		}
		bioRecord.RecordedTime, err = parseDateTime(date, timeStr, location)
		if err != nil {
			return nil, fmt.Errorf("parsing biometric time: %w", err)
		}
		records = append(records, bioRecord)
	}

	return records, nil

}

func parseNutrientFloat(value, nutrient string) (float64, error) {
	f, err := parseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing %s value %q: %w", nutrient, value, err)
	}
	return f, nil
}
