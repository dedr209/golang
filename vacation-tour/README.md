# Vacation Tour Calculator

Separate project for the travel agency task.

## Pricing rules

- Bulgaria: `summer $100`, `winter $150`
- Germany: `summer $160`, `winter $200`
- Poland: `summer $120`, `winter $180`
- Guide fee: `$50` per day
- Luxury room: `+20%` to destination seasonal price before guide fee is added

Formula:

`total = seasonal_destination_price * room_multiplier + (days * 50)`

Where `room_multiplier = 1.2` for `luxury`, otherwise `1.0`.

## Run

```bash
go run .
```

## Test

```bash
go test ./...
```
