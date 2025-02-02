package matrix

import (
	"matrices/matrixtools"
)

// Normalize erwartet eine Spaltennummer.
// Falls das Diagonalelement [col][col] nicht 0 ist, wird die Zeile durch das Diagonalelement normiert.
// D.h. die gesamte Zeile col wird durch das Diagonalelement geteilt.
func (m Matrix) Normalize(col int) {
	factor := m[col][col]
	matrixtools.ScalarMultRow(m, col, 1/factor)
}

// EliminateBelow erwartet eine Spaltennummer `col`.
// Multipliziert alle Zeilen row unter der Zeile col mit -1/Matrix[row][col] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element unter dem Diagonalelement 0.
// Voraussetzung: Die Zeile col ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateBelow(col int) {

	for i := col + 1; i < len(m); i++ {
		factor := m[i][col]
		matrixtools.ScalarMultRow(m, i, 1/factor)
		matrixtools.SubRows(m, i, col)
	}

}

// EliminateAbove erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen über der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element über dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateAbove(row int) {
	for i := row - 1; i > -1; i-- {
		factor := m[i][row]
		matrixtools.ScalarMultRow(m, i, 1/factor)
		matrixtools.SubRows(m, i, row)
	}
}

// UpperTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine obere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) UpperTriangular() {
	for i := range m {
		m.Normalize(i)
		m.EliminateBelow(i)

	}
}

// LowerTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine untere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) LowerTriangular() {
	for i := len(m); i > 0; i-- {
		m.Normalize(i - 1)
		m.EliminateAbove(i - 1)
	}
}

// Gauss transformiert die Matrix im linken Bereich in die Einheitsmatrix.
func (m Matrix) Gauss() {
	m.UpperTriangular()
	m.LowerTriangular()
}
