# Coincheck

## Coincheck is a simple command line program written for quickly finding crytocurrency prices.

This is the primary version of Coincheck. It is a port written in Go of an older, unmaintained version (see
<https://github.com/ayuopy/coincheck-old>).

Once cloned, you can copy the binary to `usr/local/bin` or `$HOME/bin` to access it from all directories.

## Usage

`coincheck [-c] [COINS] ...`


## Options

```
-c      specify a 3 character currency code (e.g. gbp or eur) to use in place of USD
-h      get help text
```

## Examples


Fetch latest Bitcoin and Litecoin prices with case-insensitive symbol or name:

```
$ coincheck btc litecoin
--------------------------------------------------------------------------------
Rank  Symbol  Name              USD Price   BTC Price     24h Change   7d Change
--------------------------------------------------------------------------------
1     BTC     Bitcoin           2630.19     1.0           -3.88%       -6.49%
7     LTC     Litecoin          30.26       0.0115921     1.41%        3.1%
--------------------------------------------------------------------------------
```

Fetch top 10 coins in GBP:

```
$ coincheck -c gbp
--------------------------------------------------------------------------------
Rank  Symbol  Name              GBP Price   BTC Price     24h Change   7d Change
--------------------------------------------------------------------------------
1     BTC     Bitcoin           2065.09     1.0           -3.88%       -6.49%
2     ETH     Ethereum          296.00      0.144443      -3.36%       44.32%
3     XRP     Ripple            0.22        0.00010835    10.05%       -2.17%
4     XEM     NEM               0.16        0.00008040    0.88%        -6.71%
5     ETC     Ethereum Classic  15.31       0.00746889    -3.36%       8.79%
6     IOT     IOTA              0.45        0.00021896    -10.05%      %
7     LTC     Litecoin          23.76       0.0115921     1.41%        3.1%
8     DASH    Dash              134.99      0.0658722     -4.17%       14.16%
9     BTS     BitShares         0.26        0.00012849    -7.24%       196.59%
10    STRAT   Stratis           6.78        0.00330785    0.36%        -14.36%
--------------------------------------------------------------------------------
```

## Future Updates

* Port feature from coincheck-old showing how much a currency has in/decreased in
    price since last searched for.
* Possibly change method used to find currency information as the Coin Market Cap API
    can be slow to call.
