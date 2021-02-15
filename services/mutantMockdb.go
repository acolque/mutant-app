package services

import "errors"

type MutantMockDb struct {
	dnas []EDna
}

func NewMutantMockDb() *MutantMockDb {
	return &MutantMockDb{dnas: make([]EDna, 0)}
}

func (db *MutantMockDb) Find(dna EDna) (EDna, error) {
	var result EDna
	for _, d := range db.dnas {
		if compare(d.Dna, dna.Dna) {
			return d, nil
		}
	}
	return result, errors.New("Element not exist")
}

func (db *MutantMockDb) Add(dna EDna) error {
	db.dnas = append(db.dnas, dna)
	return nil
}

func (db *MutantMockDb) Update(dna EDna, newDna EDna) error {
	return nil
}

func (db *MutantMockDb) Delete(dna EDna) error {
	for i, d := range db.dnas {
		if compare(d.Dna, dna.Dna) {
			db.dnas = append(db.dnas[:i], db.dnas[i+1:]...)
			break
		}
	}
	return nil
}

func (db *MutantMockDb) GetAll() ([]EDna, error) {
	return db.dnas, nil
}

func compare(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i, v := range arr1 {
		if v != arr2[i] {
			return false
		}
	}

	return true
}
