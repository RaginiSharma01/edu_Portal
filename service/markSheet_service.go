package service

import (
	"bytes"
	"context"
	"fmt"
	"smp/models"
	"smp/repository"
	"sort"

	"github.com/jung-kurt/gofpdf"
)

type MarksService struct {
	repo *repository.MarksRepository
}

func NewMarksService(repo *repository.MarksRepository) *MarksService {
	return &MarksService{repo: repo}
}

func (s *MarksService) CreateMarks(ctx context.Context, m models.CreateMarks) error {
	return s.repo.CreateMarks(ctx, m)
}

func (s *MarksService) GetMarks(ctx context.Context, term string) ([]models.StudentMarks, error) {
	return s.repo.GetMarks(ctx, term)
}

// ✅ Single student PDF — fully dynamic subjects
func (s *MarksService) GenerateStudentPDF(ctx context.Context, studentID, term string) ([]byte, error) {
	m, err := s.repo.GetMarksByStudentID(ctx, studentID, term)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// Title
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(190, 10, "Student Marksheet", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// Student Info
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(190, 8, "Name: "+m.Student)
	pdf.Ln(6)
	pdf.Cell(190, 8, "Term: "+term)
	pdf.Ln(10)

	// Table Header
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(95, 10, "Subject", "1", 0, "C", false, 0, "")
	pdf.CellFormat(95, 10, "Marks", "1", 1, "C", false, 0, "")

	// ✅ Sort subjects alphabetically for consistent order
	subjectNames := make([]string, 0, len(m.Subjects))
	for name := range m.Subjects {
		subjectNames = append(subjectNames, name)
	}
	sort.Strings(subjectNames)

	// Table Rows — dynamic
	pdf.SetFont("Arial", "", 12)
	maxPerSubject := 0
	if len(subjectNames) > 0 {
		maxPerSubject = m.MaxTotal / len(subjectNames)
	}

	for _, subj := range subjectNames {
		pdf.CellFormat(95, 10, subj, "1", 0, "", false, 0, "")
		pdf.CellFormat(95, 10, fmt.Sprintf("%d / %d", m.Subjects[subj], maxPerSubject), "1", 1, "C", false, 0, "")
	}

	pdf.Ln(8)

	// Summary
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(95, 10, fmt.Sprintf("Total: %d / %d", m.Total, m.MaxTotal), "1", 0, "", false, 0, "")
	pdf.CellFormat(95, 10, fmt.Sprintf("Percentage: %.2f%%", m.Percentage), "1", 1, "", false, 0, "")

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// ✅ All students PDF — fully dynamic subjects
func (s *MarksService) GenerateMarksPDF(ctx context.Context, term string) ([]byte, error) {
	marks, err := s.repo.GetMarks(ctx, term)
	if err != nil {
		return nil, err
	}

	if len(marks) == 0 {
		return nil, fmt.Errorf("no marks found for term %s", term)
	}

	// ✅ Collect all unique subject names across all students
	subjectSet := make(map[string]struct{})
	for _, m := range marks {
		for subj := range m.Subjects {
			subjectSet[subj] = struct{}{}
		}
	}
	subjectNames := make([]string, 0, len(subjectSet))
	for name := range subjectSet {
		subjectNames = append(subjectNames, name)
	}
	sort.Strings(subjectNames)

	pdf := gofpdf.New("L", "mm", "A4", "") // Landscape
	pdf.AddPage()

	// Title
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(277, 10, fmt.Sprintf("Marksheet - Term %s", term), "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// ✅ Dynamic column widths
	nameColW := 50.0
	subjectColW := 25.0
	totalColW := 30.0
	pctColW := 27.0

	// Header row
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)
	pdf.CellFormat(nameColW, 10, "Student", "1", 0, "C", true, 0, "")
	for _, subj := range subjectNames {
		pdf.CellFormat(subjectColW, 10, subj, "1", 0, "C", true, 0, "")
	}
	pdf.CellFormat(totalColW, 10, "Total", "1", 0, "C", true, 0, "")
	pdf.CellFormat(pctColW, 10, "%", "1", 1, "C", true, 0, "")

	// Data rows
	pdf.SetFont("Arial", "", 10)
	for _, m := range marks {
		pdf.CellFormat(nameColW, 10, m.Student, "1", 0, "L", false, 0, "")
		for _, subj := range subjectNames {
			val := m.Subjects[subj] // 0 if subject not present for this student
			pdf.CellFormat(subjectColW, 10, fmt.Sprintf("%d", val), "1", 0, "C", false, 0, "")
		}
		pdf.CellFormat(totalColW, 10, fmt.Sprintf("%d / %d", m.Total, m.MaxTotal), "1", 0, "C", false, 0, "")
		pdf.CellFormat(pctColW, 10, fmt.Sprintf("%.2f%%", m.Percentage), "1", 1, "C", false, 0, "")
	}

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
