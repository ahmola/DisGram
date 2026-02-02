# -g : General API 정보가 있는 파일 (main 함수가 있는 곳)
# -d : 스캔할 디렉토리 (콤마로 구분). 여기서는 main과 internal 폴더를 다 봐야 함
# --output : 문서가 생성될 위치
# MUST BE EXECUTED IN EACH SERVICE DIRECTORIES!!!
swag init -g main.go -d main,internal --output main/docs --parseDependency --parseInternal