@echo off
setlocal

echo Generating protobuf files...

if not exist gen mkdir gen
if not exist gen\apeiron mkdir gen\apeiron
if not exist gen\apeiron\v1 mkdir gen\apeiron\v1

protoc ^
  --proto_path=proto ^
  --go_out=gen ^
  --go_opt=paths=source_relative ^
  --go-grpc_out=gen ^
  --go-grpc_opt=paths=source_relative ^
  proto/apeiron/v1/common.proto ^
  proto/apeiron/v1/cache_service.proto ^
  proto/apeiron/v1/creature_data_service.proto ^
  proto/apeiron/v1/inventory_data_service.proto ^
  proto/apeiron/v1/player_data_service.proto ^
  proto/apeiron/v1/profile_data_service.proto ^
  proto/apeiron/v1/skill_data_service.proto ^
  proto/apeiron/v1/world_data_service.proto

if %ERRORLEVEL% neq 0 (
  echo Proto generation failed.
  exit /b %ERRORLEVEL%
)

echo Proto generation completed.
endlocal
