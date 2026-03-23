#Funciona de la siguiente manera
1. Descargar proyecto
2. ejecutar go run ./cmd/server/main.go
3. en otra terminal ejecutar la url a acortar. Ejemplo:
   curl -X POST http://localhost:8080/shorten \
   -H "Content-Type: application/json" \
   -d '{"url":"https://google.com"}'
4. Se te va a devolver una url acortada
