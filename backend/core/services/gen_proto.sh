for dir in pkg/proto/*/ ; do
    protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        ${dir}*.proto
done

echo "✅ Proto generation completed!"