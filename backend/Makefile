sqlboiler:
	sqlboiler mysql

mock_usecase:
	mockgen -source=internal/usecase/usecase.go -destination=internal/usecase/mock/usecase_mock.go -package=mock

mock_repo:
	mockgen -source=internal/interface/repository/repository.go -destination=internal/interface/repository/mock/repository_mock.go -package=mock

go_test:
	go test -v -cover ./internal/usecase -coverprofile=cover.out
	go tool cover -html=cover.out -o=cover.html
	open cover.html

go_test_rm:
	rm cover.out cover.html
