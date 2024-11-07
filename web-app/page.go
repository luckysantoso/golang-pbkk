package main

import (
	"os"
)

// Page struct untuk menyimpan data judul dan isi halaman
type Page struct {
    Title string
    Body  []byte
}

// Fungsi untuk menyimpan halaman ke file dengan nama sesuai judul
func (p *Page) save() error {
    filename := p.Title + ".txt"
    return os.WriteFile(filename, p.Body, 0600) // Menyimpan konten halaman dengan WriteFile dari paket os
}

// Fungsi untuk memuat halaman dari file berdasarkan judul
func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := os.ReadFile(filename) // Membaca file dengan ReadFile dari paket os
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}
