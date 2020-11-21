# trimfaketime

trimfaketime is a filter that trims Go Fake Time playback headers of streams.

## Install

```
go get -u github.com/knightso/trimfaketime/cmd/tft
```

## Usage

```
go test --tags=faketime . | tft
```

Note: Go test command status will be lost using pipes like above. You can get it by some shell features like PIPESTATUS variable if you need.
