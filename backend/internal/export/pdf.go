package export

import (
    "cenova_nabidka/backend/internal/offer"
    "github.com/jung-kurt/gofpdf"
)

func GeneratePDF(offer *offer.PriceOffer) ([]byte, error) {
    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()

    pdf.SetFont("Arial", "B", 16)
    pdf.Cell(40, 10, "Cenová Nabídka")

    pdf.Ln(12)

    pdf.SetFont("Arial", "", 12)
    pdf.Cell(40, 10, "Zákazník:")
    pdf.Ln(10)
    pdf.Cell(40, 10, offer.FirstName+" "+offer.LastName)
    pdf.Ln(6)
    pdf.Cell(40, 10, offer.Phone)
    pdf.Ln(6)
    pdf.Cell(40, 10, offer.Email)
    pdf.Ln(6)
    pdf.Cell(40, 10, offer.City)

    pdf.Ln(12)

    pdf.SetFont("Arial", "B", 12)
    pdf.CellFormat(40, 10, "Produkt", "1", 0, "C", false, 0, "")
    pdf.CellFormat(30, 10, "Cena", "1", 0, "C", false, 0, "")
    pdf.CellFormat(30, 10, "Sleva", "1", 0, "C", false, 0, "")
    pdf.CellFormat(30, 10, "Celkem", "1", 0, "C", false, 0, "")
    pdf.Ln(-1)

    pdf.SetFont("Arial", "", 12)
    for _, poProduct := range offer.Products {
        pdf.CellFormat(40, 10, poProduct.Product.Name, "1", 0, "", false, 0, poProduct.Product.URL)
        pdf.CellFormat(30, 10, fmt.Sprintf("%.2f Kč", poProduct.UnitPrice), "1", 0, "R", false, 0, "")
        pdf.CellFormat(30, 10, fmt.Sprintf("%.2f Kč", poProduct.Product.DiscountSW), "1", 0, "R", false, 0, "")
        pdf.CellFormat(30, 10, fmt.Sprintf("%.2f Kč", poProduct.TotalPrice), "1", 0, "R", false, 0, "")
        pdf.Ln(-1)
    }

    // Výpočet celkové ceny
    pdf.SetFont("Arial", "B", 12)
    pdf.CellFormat(70, 10, "Celková Cena:", "1", 0, "R", false, 0, "")
    pdf.CellFormat(30, 10, fmt.Sprintf("%.2f Kč", offer.TotalPrice), "1", 0, "R", false, 0, "")
    pdf.Ln(-1)

    // Vygenerování PDF do byte array
    var buf bytes.Buffer
    err := pdf.Output(&buf)
    if err != nil {
        return nil, err
    }

    return buf.Bytes(), nil
}
