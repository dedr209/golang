package bank

import "testing"

func TestNewClientAndAccessors(t *testing.T) {
	client := NewClient("John", "Doe", "AC-1", 1000, 250)

	if client.GetName() != "John" {
		t.Fatalf("expected name John, got %s", client.GetName())
	}
	if client.GetSurname() != "Doe" {
		t.Fatalf("expected surname Doe, got %s", client.GetSurname())
	}
	if client.GetAccountNumber() != "AC-1" {
		t.Fatalf("expected account AC-1, got %s", client.GetAccountNumber())
	}
	if client.GetDeposit() != 1000 {
		t.Fatalf("expected deposit 1000, got %f", client.GetDeposit())
	}
	if client.GetCredit() != 250 {
		t.Fatalf("expected credit 250, got %f", client.GetCredit())
	}

	client.SetName("Jane")
	client.SetSurname("Smith")
	client.SetAccountNumber("AC-2")
	client.SetDeposit(1200)
	client.SetCredit(500)

	if client.GetName() != "Jane" || client.GetSurname() != "Smith" || client.GetAccountNumber() != "AC-2" {
		t.Fatalf("client string setters were not applied")
	}
	if client.GetDeposit() != 1200 || client.GetCredit() != 500 {
		t.Fatalf("client money setters were not applied")
	}
}

func TestNewBankAndClientManagement(t *testing.T) {
	b := NewBank("Test Bank", 50000, 7000, 3000)

	if b.GetName() != "Test Bank" {
		t.Fatalf("expected bank name Test Bank, got %s", b.GetName())
	}
	if b.GetBankMoney() != 50000 || b.GetDeposit() != 7000 || b.GetCredit() != 3000 {
		t.Fatalf("bank constructor values are incorrect")
	}

	b.SetName("Updated Bank")
	b.SetBankMoney(75000)
	b.SetDeposit(8200)
	b.SetCredit(4300)

	if b.GetName() != "Updated Bank" || b.GetBankMoney() != 75000 || b.GetDeposit() != 8200 || b.GetCredit() != 4300 {
		t.Fatalf("bank setters were not applied")
	}

	client := NewClient("Alice", "Brown", "AC-9", 2500, 900)
	b.AddClient(client)

	stored, ok := b.GetClient("AC-9")
	if !ok {
		t.Fatalf("expected client AC-9 to exist")
	}
	if stored.GetName() != "Alice" {
		t.Fatalf("expected client name Alice, got %s", stored.GetName())
	}

	allClients := b.GetClients()
	delete(allClients, "AC-9") // should not mutate bank internals

	if _, ok := b.GetClient("AC-9"); !ok {
		t.Fatalf("defensive copy failed, bank data was mutated")
	}

	b.RemoveClient("AC-9")
	if _, ok := b.GetClient("AC-9"); ok {
		t.Fatalf("client should be removed")
	}
}
