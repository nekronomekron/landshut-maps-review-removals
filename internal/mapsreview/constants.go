package mapsreview

const (
	OutputDir     = "output"
	ResultsJSON   = "output/places.json"
	ResultsCSV    = "output/places.csv"
	DiscoveryJSON = "output/discovery.json"
	MetadataJSON  = "output/metadata.json"
)

var NurembergPostcodes = []string{
	"84028",
	"84030",
	"84032",
	"84034",
	"84036",
	"84144",
	"84137",
}

var DefaultQueries = []string{
	// Gastro (original)
	"restaurant", "café", "imbiss", "pizzeria", "bäckerei",
	"döner", "burger", "sushi", "schnitzel", "frühstück", "brunch",
	// Bars & Nightlife
	"bar", "kneipe", "pub", "biergarten", "brauerei",
	"cocktail bar", "lounge", "weinstube",
	"club", "nachtclub", "diskothek",
	// Hotels
	"hotel",
	// Beauty & Wellness
	"friseur", "barbier", "barbershop",
	"fitnessstudio", "fitness",
	// Shopping & Daily
	"supermarkt", "metzgerei",
	"apotheke",
	// Services
	"tankstelle",
}

var NurembergPostcodeSet = func() map[string]bool {
	set := make(map[string]bool, len(NurembergPostcodes))
	for _, postcode := range NurembergPostcodes {
		set[postcode] = true
	}
	return set
}()
