package resistorcolor

var colors = []string{
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}


func Colors() []string {
	return colors 
}


func ColorCode(color string) int {
    var res int 
	for i , c := range colors {
        if c == color {
            res = i
        }
    }
    return res 
}
