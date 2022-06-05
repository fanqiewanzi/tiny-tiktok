userservice:
	kitex -I idl/ -module github.com/weirdo0314/tiny-tiktok -type protobuf idl/user.proto
videoservice:
	kitex -I idl/ -module github.com/weirdo0314/tiny-tiktok -type protobuf idl/video.proto