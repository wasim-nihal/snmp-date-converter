# SNMP DateAndTime Hex Converter

This is a simple go unility program that converts a hex-encoded SNMP `DateAndTime` value into a human-readable date and time string.

It was written for personal debugging purposes, but anyone is free to use if needed and this is not maintained.

## What It Does

SNMP (Simple Network Management Protocol) defines a `DateAndTime` format using a specific byte structure. This tool takes the hex representation of that format and converts it into a readable timestamp.

The SNMP format can be found in the RFC: https://datatracker.ietf.org/doc/html/rfc2579

For example, input like:

```
07 E9 07 0A 0E 1E 2D 00 2B 02 00
```

Will output:

```
Parsed SNMP DateTime:
2025-07-10 14:30:45.0 +02:00
```

## SNMP DateAndTime Format

The SNMP `DateAndTime` format consists of:

| Byte Index | Meaning                 |
|------------|-------------------------|
| 0–1        | Year (2 bytes)          |
| 2          | Month (1–12)            |
| 3          | Day (1–31)              |
| 4          | Hour (0–23)             |
| 5          | Minute (0–59)           |
| 6          | Second (0–60)           |
| 7          | Deci-seconds (0–9)      |
| 8          | Direction from UTC (`+` or `-`) |
| 9          | UTC Offset Hours (0–13) |
| 10         | UTC Offset Minutes (0–59) |

## Usage

### Build

```bash
go build -o snmp-date-converter
```

### Run

```bash
./snmp-date-converter -convert "07 E9 07 0A 0E 1E 2D 00 2B 02 00"
```

### Output

```bash
Parsed SNMP DateTime:
2025-07-10 14:30:45.0 +02:00
```

## Input Format

- The input should be a space-separated hex string of exactly 11 bytes (22 hex characters).
- The 9th byte (UTC direction) must be either `2B` (`+`) or `2D` (`-`).

## Flags

| Flag         | Description                                                   |
|--------------|---------------------------------------------------------------|
| `-convert`   | The hex string input (required). Example: `07 E9 07 0A ...`   |

## Example Conversion

Input:
```
07 E9 07 0A 0E 1E 2D 00 2B 02 00
```

Parsed:
- Year: 2025
- Month: 07 (July)
- Day: 10
- Time: 14:30:45.0
- Timezone: +02:00

Output:
```
2025-07-10 14:30:45.0 +02:00
```

