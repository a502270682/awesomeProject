package tools

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func getHeader() []string {
	return []string{"Product", "Quantity", "Price"}
}

func getListContent() [][]string {
	list := make([][]string, 0)
	return list
}

func getContents() [][]string {
	return [][]string{
		{"Swamp", "12", "R$ 4,00"},
		{"Sorin, A Planeswalker", "4", "R$ 90,00"},
		{"Tassa", "4", "R$ 30,00"},
		{"Skinrender", "4", "R$ 9,00"},
		{"Island", "12", "R$ 4,00"},
		{"Mountain", "12", "R$ 4,00"},
		{"Plain", "12", "R$ 4,00"},
		{"Black Lotus", "1", "R$ 1.000,00"},
		{"Time Walk", "1", "R$ 1.000,00"},
		{"Emberclave", "4", "R$ 44,00"},
		{"Anax", "4", "R$ 32,00"},
	}
}

func getDarkGrayColor() color.Color {
	return color.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() color.Color {
	return color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() color.Color {
	return color.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() color.Color {
	return color.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}

func TestPdfTable(t *testing.T) {
	begin := time.Now()

	//darkGrayColor := getDarkGrayColor()
	grayColor := getGrayColor()
	//whiteColor := color.NewWhite()
	header := getHeader()
	contents := getContents()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
	m.SetBorder(true)

	//m.RegisterHeader(func() {
	//	m.Row(20, func() {
	//		//m.Col(3, func() {
	//		//	_ = m.FileImage("internal/assets/images/biplane.jpg", props.Rect{
	//		//		Center:  true,
	//		//		Percent: 80,
	//		//	})
	//		//})
	//
	//		//m.ColSpace(6)
	//
	//		m.Col(3, func() {
	//			m.Text("AnyCompany Name Inc. 851 Any Street Name, Suite 120, Any City, CA 45123.", props.Text{
	//				Size:        8,
	//				Align:       consts.Right,
	//				Extrapolate: false,
	//			})
	//			m.Text("Tel: 55 024 12345-1234", props.Text{
	//				Top:   12,
	//				Style: consts.BoldItalic,
	//				Size:  8,
	//				Align: consts.Right,
	//			})
	//			m.Text("www.mycompany.com", props.Text{
	//				Top:   15,
	//				Style: consts.BoldItalic,
	//				Size:  8,
	//				Align: consts.Right,
	//			})
	//		})
	//	})
	//})
	//
	//m.RegisterFooter(func() {
	//	m.Row(20, func() {
	//		m.Col(12, func() {
	//			m.Text("Tel: 55 024 12345-1234", props.Text{
	//				Top:   13,
	//				Style: consts.BoldItalic,
	//				Size:  8,
	//				Align: consts.Left,
	//			})
	//			m.Text("www.mycompany.com", props.Text{
	//				Top:   16,
	//				Style: consts.BoldItalic,
	//				Size:  8,
	//				Align: consts.Left,
	//			})
	//		})
	//	})
	//})
	//
	//m.Row(10, func() {
	//	m.Col(12, func() {
	//		m.Text("Invoice ABC123456789", props.Text{
	//			Top:   3,
	//			Style: consts.Bold,
	//			Align: consts.Center,
	//		})
	//	})
	//})
	//
	//m.SetBackgroundColor(darkGrayColor)
	//
	//m.Row(7, func() {
	//	m.Col(3, func() {
	//		m.Text("Transactions", props.Text{
	//			Top:   1.5,
	//			Size:  9,
	//			Style: consts.Bold,
	//			Align: consts.Center,
	//		})
	//	})
	//	m.ColSpace(9)
	//})
	//
	//m.SetBackgroundColor(whiteColor)

	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{4, 2, 3},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{4, 2, 3},
		},
		Align:                consts.Center,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	err := m.OutputFileAndClose("./billing.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func TestCreatePdf1(t *testing.T) {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	darkGrayColor := getDarkGrayColor()
	whiteColor := color.NewWhite()

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(3, func() {
				_ = m.FileImage("/Users/guofeiyang/Downloads/111.png", props.Rect{
					Center:  true,
					Percent: 80,
				})
			})

			m.ColSpace(6)

			m.Col(3, func() {
				m.Text("AnyCompany Name Inc. 851 Any Street Name, Suite 120, Any City, CA 45123.", props.Text{
					Size:        8,
					Align:       consts.Right,
					Extrapolate: false,
				})
				m.Text("Tel: 55 024 12345-1234", props.Text{
					Top:   12,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
				})
				m.Text("www.mycompany.com", props.Text{
					Top:   15,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Right,
				})
			})
		})
	})
	m.SetPageMargins(10, 15, 10)
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Invoice ABC123456789", props.Text{
				Top:   3,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
	})

	m.SetBackgroundColor(darkGrayColor)

	m.Row(7, func() {
		m.Col(3, func() {
			m.Text("Transactions", props.Text{
				Top:   1.5,
				Size:  9,
				Style: consts.Bold,
				Align: consts.Center,
			})
		})
		m.ColSpace(9)
	})

	m.SetBackgroundColor(whiteColor)
	err := m.OutputFileAndClose("./billing.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
}

func TestMergePdf(t *testing.T) {
	//darkGrayColor := getDarkGrayColor()
	grayColor := getGrayColor()
	//whiteColor := color.NewWhite()
	header := getHeader()
	contents := getContents()

	m := pdf.NewMaroto(consts.Portrait, consts.A4)

	m.SetBorder(true)
	m.TableList(header, contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: []uint{4, 2, 3},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{4, 2, 3},
		},
		Align:                consts.Center,
		AlternatedBackground: &grayColor,
		HeaderContentSpace:   1,
		Line:                 false,
	})

	err := m.OutputFileAndClose("test.pdf")
	if err != nil {
		t.Fatal(err)
	}

	//inFile := []string{filepath.Join(".", "billing.pdf"), filepath.Join(".", "test.pdf")}
	inFile := []string{filepath.Join(".", "billing.pdf"), filepath.Join(".", "test.pdf")}

	err = api.MergeAppendFile(inFile, "merge.pdf", nil)
	if err != nil {
		t.Fatal(err)
	}

}
