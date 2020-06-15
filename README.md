# days
A simple program for calculating days difference between dates.

## Usage
```
Usage:
  days [command]

Available Commands:
  add         Date in a given days after given date. You can omit date, today's date will be taken.
  between     amount of days between two dates
  help        Help about any command
  rm          date in a given days before given date. You can omit date, today's date will be taken.
  to          days between now and future date
  weekday     prints the name of the day at given date

Flags:
  -h, --help   help for days

Use "days [command] --help" for more information about a command.
```

## Example
```
>> go install . 
>> days weekday 27-10-1995 21-10-2020 16-05-2020
27 Oct 1995 was Friday
21 Oct 2020 will be Wednesday
16 May 2020 was Saturday

>> days add --help
Date in a given days after given date. You can omit date, today's date will be taken.

Usage:
 days add [days: number] [date: yyyy-mm-dd] [flags]

Flags:
 -h, --help   help for add


>> days add 10000 27-10-1995
Tue 14 Mar 2023

>> days between 27-10-1995 14-03-2023
Fri 27 Oct 1995 is 10000 days before Tue 14 Mar 2023.

>> days between 14-03-2023 27-10-1995
Tue 14 Mar 2023 is 10000 days after Fri 27 Oct 1995.

```

