package pkg

import "errors"

var New = errors.New

var (
	ErrFormatRequestBody = errors.New("Format data yang dikirim tidak valid.")
	ErrInternal          = errors.New("Terdapat kesalahan pada internal")
	ErrMethodNotAllow    = errors.New("Metode tidak diizinkan")
	ErrInvalidURL        = errors.New("URL tidak valid")
	ErrGetDataRedis      = errors.New("Gagal mengambil redis")
)
