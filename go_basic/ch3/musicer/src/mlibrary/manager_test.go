package mlibrary

import (
	"testing"
)

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManger failed.")
	}

	if mm.Len() != 0 {
		t.Error("NewMusicManger failed, not empty.")
	}

	m0 := &MusicEntry{
		"1", "My Heart Will Go On", "Celion Dion", Pop, "http://qbox.me//24501234", MP3}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("NewMusicManger.Add() failed.")
	}

	m :=mm.Find(m0.Name) {
		t.Error("NewMusicManger.Find() failed.")
	}

	if m.Id != m0.Id || m.Artist != m0.Artist || 
		m.Name != m0.Name || m.Genre != m0.Genre ||
		m.Source != m0.Source || m.Type != m0.Type {
			t.Error("NewMusicManger.Find failed. Found item mismatch.")
		}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("NewMusicManger.Get() failed.", err)
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("NewMusicManger.Remove() failed", err)
	}
}