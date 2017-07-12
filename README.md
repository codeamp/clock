# Travel through time with clock

```
clock.NowFunc = func() time.Time {
  return time.Date(2017, 07, 11, 18, 20, 00, 0, time.UTC)
}
```

instead of `time.Now()` you can use `clock.Now()`
