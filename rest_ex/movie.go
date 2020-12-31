package main

/* Modelo */
type Movie struct{
  Name string     `json:"name"`
  Year int        `json:"year"`
  Director string `json:"director"`
}
/* Objeto generalizado / Slice */
type Movies []Movie
