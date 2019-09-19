protogen:
	./scripts/protogen.sh

run:
	reflex -c reflex.conf
	
dist:
	@./scripts/dist.sh

test:
	ginkgo -r -v -coverpkg ./v1/...

reflex:
	reflex -r '(\.go|go\.mod)' -- ginkgo -r -v -coverpkg ./v1/...