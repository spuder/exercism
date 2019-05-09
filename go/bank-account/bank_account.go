package account


type Account struct {
	balance int64
	closed  bool
}

func Open(initialDeposit int64) *Account  {
  if initialDeposit >= 0 {
    a := new(Account)
    a.balance = initialDeposit
    a.closed = false
    return a
  } else {
    return nil
  }

}

func (a *Account ) Balance() (balance int64, ok bool) {
  if a.closed {
    ok = false
  } else {
    ok = true
  }
  return a.balance, ok
}

func (a *Account) Close() (payout int64, ok bool) {
  if a.closed {
    ok = false
  } else {
    payout = a.balance
    a.balance = 0
    a.closed = true
    ok = true
  }
  return payout, ok
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
  if a.closed == true {
    ok = false
  } else {
    if amount < 0 {
      if a.balance + amount < 0 {
        ok = false
      } else {
        a.balance = a.balance + amount
        ok = true
      }
    } else {
      a.balance += amount
      ok = true
    }

  }
  return a.balance, ok
}