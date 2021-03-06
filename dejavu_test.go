package dejavu

import "testing"

func TestDeterministic(t *testing.T) {

	d := NewDeterministic(3)

	// add entries
	if d.Witness([]byte("foo")) {
		t.Errorf("Incorrect déjà vu: 'foo'!")
	}
	if d.Witness([]byte("bar")) {
		t.Errorf("Incorrect déjà vu: 'bar'!")
	}

	// remembers entry
	if !d.Witness([]byte("bar")) {
		t.Errorf("Expected déjà vu: 'bar'!")
	}

	// remembers oldest entry before overwriting
	if !d.Witness([]byte("foo")) {
		t.Errorf("Expected déjà vu: 'foo'!")
	}

	// add entries
	if d.Witness([]byte("bam")) {
		t.Errorf("Incorrect déjà vu: 'bam'!")
	}
	if d.Witness([]byte("baz")) {
		t.Errorf("Incorrect déjà vu: 'baz'!")
	}

	// forgot oldest
	if d.Witness([]byte("bar")) {
		t.Errorf("Incorrect déjà vu: 'bar'!")
	}
}

func TestProbabilistic(t *testing.T) {

	d := NewProbabilistic(1024, 0.000001)

	// add entries
	if d.Witness([]byte("foo")) {
		t.Errorf("Incorrect déjà vu: 'foo'!")
	}
	if d.Witness([]byte("bar")) {
		t.Errorf("Incorrect déjà vu: 'bar'!")
	}

	// remembers entry
	if !d.Witness([]byte("foo")) {
		t.Errorf("Expected déjà vu: 'foo'!")
	}

	// fill memory
	for i := 0; i < 2048; i++ {
		d.Witness([]byte("data"))
	}

	// forgot oldest
	if d.Witness([]byte("bar")) {
		t.Errorf("Incorrect déjà vu: 'bar'!")
	}
}
