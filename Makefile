proto:
	@protoc -I protos protos/*.proto --go_out=.