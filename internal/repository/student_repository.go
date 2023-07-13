package repository

import (
	Entity "github.com/<TEMPLATE>/internal/entity/person"
	"github.com/<TEMPLATE>/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentRepository struct {
	collection string
	Mongo      *mongo.Client
	Log        *logger.Logger
}

func NewStudentRepository(mc *mongo.Client, log *logger.Logger) StudentRepository {
	return StudentRepository{
		collection: "person",
		Mongo:      mc,
		Log:        log,
	}
}

func (atar *StudentRepository) SaveNewEntity(entityDoc *Entity.StudentEntity) *Entity.StudentEntity {
	// mClient := atar.Mongo

	// dataJSON := entityDoc.ToJson()
	// req := esapi.IndexRequest{
	// 	Index:      atar.Index,
	// 	DocumentID: uuid.New().String(),
	// 	Body:       bytes.NewReader(dataJSON),
	// 	Refresh:    "true",
	// }
	// // atar.Log.Info("SaveNewEntity: " + string(dataJSON))

	// // Perform the request with the client.
	// res, err := req.Do(context.Background(), mClient)
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }
	// defer res.Body.Close()

	// if err != nil {
	// 	panic(err)
	// }

	// if res.IsError() {
	// 	log.Printf("[%s] Error indexing document ID=%d", res.Status(), 1)
	// } else {
	// 	// Deserialize the response into a map.
	// 	var r map[string]interface{}
	// 	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	// 		log.Printf("Error parsing the response body: %s", err)
	// 	} else {
	// 		// Print the response status and indexed document version.
	// 		log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))

	// 		// sj, _ := json.Marshal(r)
	// 		// log.Printf("---> %s", sj)
	// 	}
	// }

	return entityDoc
}

func makeFieldsForSearch(entity Entity.StudentEntity) map[string]interface{} {
	query_search := map[string]interface{}{}

	fields := entity.ToMap()

	for key, value := range fields {
		query_search[key] = map[string]interface{}{
			"query": value,
		}
	}

	return map[string]interface{}{
		"query": map[string]interface{}{
			"match_bool_prefix": query_search,
		},
	}
}

func (atar *StudentRepository) Search(fields Entity.StudentEntity) []*Entity.StudentEntity {
	// es := atar.EsClient

	// var buf bytes.Buffer
	// query := makeFieldsForSearch(fields)
	// if err := json.NewEncoder(&buf).Encode(query); err != nil {
	// 	log.Fatalf("Error encoding query: %s", err)
	// }

	// // atar.Log.Info("Map Entity: ", fields, fields.ToMap())
	// // atar.Log.Info("Search Query: " + buf.String())

	// // Perform the search request.
	// res, err := es.Search(
	// 	es.Search.WithContext(context.Background()),
	// 	es.Search.WithIndex(atar.Index),
	// 	es.Search.WithBody(&buf),
	// 	es.Search.WithTrackTotalHits(true),
	// 	es.Search.WithPretty(),
	// )
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }
	// defer res.Body.Close()

	// if res.IsError() {
	// 	var e map[string]interface{}
	// 	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
	// 		log.Fatalf("Error parsing the response body: %s", err)
	// 	} else {
	// 		// Print the response status and error information.
	// 		log.Fatalf("[%s] %s: %s",
	// 			res.Status(),
	// 			e["error"].(map[string]interface{})["type"],
	// 			e["error"].(map[string]interface{})["reason"],
	// 		)
	// 	}
	// }

	// var r map[string]interface{}

	// if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	// 	log.Fatalf("Error parsing the response body: %s", err)
	// }
	// // Print the response status, number of results, and request duration.
	// // log.Printf(
	// // 	"[%s] %d hits; took: %dms",
	// // 	res.Status(),
	// // 	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	// // 	int(r["took"].(float64)),
	// // )

	// // Print the ID and document source for each hit.
	var result []*Entity.StudentEntity
	// for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	// 	// log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	// 	result = append(result, Entity.NewAtaEntity(hit.(map[string]interface{})["_source"]))
	// }

	return result
}
