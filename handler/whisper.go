package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"voit.api/pkg/response"
)

type WhisperModel struct {
	// Alias is the human-friendly name of the model
	Alias string `json:"alias"`

	// Name is what the model should be saved as locally
	Name string `json:"name"`

	// Description is a short description of the model and its recommended use
	Description string `json:"description"`

	// DownloadURL is the URL to download the model from - this will be proxied to directly
	DownloadURL string `json:"download_url"`

	// SupportsCoreML is whether or not the model supports CoreML and has a download URL for it
	SupportsCoreML bool `json:"supports_coreml"`

	// CoreMLDownloadURL is the URL to download the CoreML model from - this will be proxied to directly
	CoreMLDownloadURL string `json:"coreml_download_url"`

	// IsDefault is whether or not the model is the default model to use
	IsDefault bool `json:"is_default"`
}

var models = []WhisperModel{
	{
		Alias:             "Tiny",
		Name:              "tiny",
		Description:       "Lightweight yet performant enough for older devices",
		DownloadURL:       "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-tiny.bin",
		SupportsCoreML:    true,
		CoreMLDownloadURL: "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-tiny.en-encoder.mlmodelc.zip",
		IsDefault:         true,
	},
	{
		Alias:             "Standard",
		Name:              "base",
		Description:       "Recommended for most users, this model is reliable and efficient on newer devices, providing high-quality transcriptions without slowing you down.",
		DownloadURL:       "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-base.bin",
		SupportsCoreML:    true,
		CoreMLDownloadURL: "https://huggingface.co/ggerganov/whisper.cpp/resolve/main/ggml-base-encoder.mlmodelc.zip",
		IsDefault:         false,
	},
}

func GetAllWhisperModels(w http.ResponseWriter, r *http.Request) {
	response.OK(w, r, &response.OKResponse{Data: models})
}

func DownloadWhisperModel(w http.ResponseWriter, r *http.Request) {
	model := chi.URLParam(r, "model")

	for _, m := range models {
		if m.Name == model {
			http.Redirect(w, r, m.DownloadURL, http.StatusTemporaryRedirect)
			return
		}
	}

	response.Error(w, r, &response.ErrResponse{
		Code:    http.StatusNotFound,
		Message: "Model not found",
	})
}

func DownloadWhisperCoreMLModel(w http.ResponseWriter, r *http.Request) {
	model := chi.URLParam(r, "model")

	for _, m := range models {
		if m.Name == model {
			http.Redirect(w, r, m.CoreMLDownloadURL, http.StatusTemporaryRedirect)
			return
		}
	}

	response.Error(w, r, &response.ErrResponse{
		Code:    http.StatusNotFound,
		Message: "Model not found",
	})
}
