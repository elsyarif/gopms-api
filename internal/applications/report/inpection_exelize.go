package report

import (
	"fmt"

	"github.com/elsyarif/pms-api/internal/domain/entities"
	"github.com/xuri/excelize/v2"
)

func InspectionExcel(data *entities.InspectionResponse) error {
	worksheet := "SINTER"
	tr := false
	scale := float64(85)

	f := excelize.NewFile()
	// create a new worksheet
	index, _ := f.NewSheet(worksheet)

	f.SetSheetView(worksheet, -1, &excelize.ViewOptions{
		DefaultGridColor: &tr,
		ShowGridLines:    &tr,
		ZoomScale:        &scale,
	})
	// Plant Name set cell
	f.MergeCell(worksheet, "B2", "E2")
	f.SetColWidth(worksheet, "B", "B", 2.55)
	f.SetColWidth(worksheet, "C", "M", 15.70)
	f.SetRowHeight(worksheet, 2, 16.45)
	f.SetCellValue(worksheet, "B2", fmt.Sprintf("Plant Name : %s", data.GroupName))
	b2, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Size:   12,
			Family: "Calibri",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "top", Color: "#000000", Style: 1},
		},
	})
	f.SetCellStyle(worksheet, "B2", "E2", b2)
	borderTop, _ := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "top", Color: "#000000", Style: 1},
		},
	})
	f.SetCellStyle(worksheet, "F2", "M2", borderTop)

	// Checker set cell
	f.MergeCell(worksheet, "B3", "F3")
	f.SetRowHeight(worksheet, 3, 16.45)
	f.SetCellValue(worksheet, "B3", fmt.Sprintf("Checker: %s", data.UserBy))
	b3, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:   true,
			Size:   12,
			Family: "Calibri",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	f.SetCellStyle(worksheet, "B3", "F3", b3)

	// Checking period set cell
	f.MergeCell(worksheet, "J3", "K3")
	f.SetCellValue(worksheet, "J3", "Checking Period :")
	f.MergeCell(worksheet, "L3", "M3")
	f.SetCellValue(worksheet, "L3", fmt.Sprintf("%s ~ %s", data.PeriodStart, data.PeriodEnd))
	j3, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "right",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   12,
			Family: "Calibri",
		},
		Border: []excelize.Border{
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	f.SetCellStyle(worksheet, "J3", "k3", j3)
	l3, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Vertical:   "center",
			Horizontal: "left",
		},
		Font: &excelize.Font{
			Bold:   true,
			Size:   12,
			Family: "Calibri",
		},
		Border: []excelize.Border{
			{Type: "right", Color: "#000000", Style: 1},
			{Type: "bottom", Color: "#000000", Style: 1},
		},
	})
	f.SetCellStyle(worksheet, "l3", "m3", l3)

	// set active worksheet to workbook
	f.SetActiveSheet(index)
	f.DeleteSheet("Sheet1")

	// Save the spreadsheet by the given path.
	err := f.SaveAs(fmt.Sprintf("%s Server Checking Weekly Resume.xlsx", data.GroupName))
	if err != nil {
		return err
	}

	return nil
}
