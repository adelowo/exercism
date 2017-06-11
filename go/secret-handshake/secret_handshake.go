package secret

const testVersion = 2

type handShake struct {
	shakeType        uint
	shakeTranslation string
}

var shakes []*handShake

func init() {
	shakes = []*handShake{
		{1, "wink"},
		{2, "double blink"},
		{4, "close your eyes"},
		{8, "jump"},
	}

}

func Handshake(secretCode uint) []string {

	var shaked []string

	shouldReverse := (secretCode & 16) > 0

	for _, shakeType := range shakes {
		if secretCode&shakeType.shakeType == 0 {
			continue
		}

		if shouldReverse {
			shaked = append([]string{shakeType.shakeTranslation}, shaked...)
			continue
		}

		shaked = append(shaked, shakeType.shakeTranslation)
	}

	return shaked
}
