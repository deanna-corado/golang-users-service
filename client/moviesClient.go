package clients

import (
	"os"

	"github.com/go-resty/resty/v2"
)

// base path for movie reqs
const moviesEndpoint = "/api/v1/movies"

// connect to movie-service
type MoviesClient struct {
	client       *resty.Client
	token        string
	clientID     string
	clientSecret string
}

func NewMoviesClient() *MoviesClient {
	return &MoviesClient{
		client:       resty.New().SetBaseURL(os.Getenv("MOVIES_SERVICE_URL")),
		clientID:     os.Getenv("MOVIES_CLIENT_ID"),
		clientSecret: os.Getenv("MOVIES_SECRET_KEY"),
	}
}

// updating token dynamically
func (mc *MoviesClient) SetToken(token string) {
	mc.token = token
}

// req with required headers
func (mc *MoviesClient) request() *resty.Request {
	return mc.client.R().
		SetHeader("Authorization", "Bearer "+mc.token)
	// SetHeader("X-Client-ID", mc.clientID).
	// SetHeader("X-Client-Secret", mc.clientSecret)
}

// GET
func (mc *MoviesClient) GetMovies() (*resty.Response, error) {
	resp, err := mc.request().Get(moviesEndpoint)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GET BY ID
func (mc *MoviesClient) GetMovieByID(movieID string) (*resty.Response, error) {
	return mc.request().Get(moviesEndpoint + "/" + movieID)
}

// POST
func (mc *MoviesClient) AddMovie(movieData map[string]any) (*resty.Response, error) {
	return mc.request().
		SetBody(movieData).
		Post(moviesEndpoint)
}

// PUT
func (mc *MoviesClient) UpdateMovie(movieID string, movieData map[string]any) (*resty.Response, error) {
	return mc.request().
		SetBody(movieData).
		Put(moviesEndpoint + "/" + movieID)
}

// DELETE
func (mc *MoviesClient) DeleteMovie(movieID string) (*resty.Response, error) {
	return mc.request().
		Delete(moviesEndpoint + "/" + movieID)
}
