type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big: big, medium: medium, small: small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big > 0 {
			this.big--
			return true
		} else {
			return false
		}
	case 2:
		if this.medium > 0 {
			this.medium--
			return true
		} else {
			return false
		}
	case 3:
		if this.small > 0 {
			this.small--
			return true
		} else {
			return false
		}
	default:
		return false
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */