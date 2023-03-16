// Code generated by gocode.Generate; DO NOT EDIT.

package filetypes

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
	_ "cuelang.org/go/pkg"
)

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenInstance.Value().FieldByName(name, true)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 1708 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xc4X\xddo\xe4\xb6\x11\x97|W\xa0R\xd3>\xe6\xad\xc0\x9c\xae\bR\xe3*#\x1f@\x81\x05\x8cC\u047b+\xee\xa5)\x8a\xf4)\b\f\xae4\xdae#\x91*I%6\xe2}h\x9b\xa6\xfd\xab\xb3\u0150\x94(je\xfb\f\xa4\xa8_\xbc;?\xcepf8\x9f\xfb\xf3\xe3\xbf\xcf\u04b3\xe3\u007f\x92\xf4\xf8\x8f$\xf9\xed\xf1\xefO\xd2\xf4=.\xb4a\xa2\xc2W\xcc0\xa2\xa7O\u04a7\u007f\x96\u04a4gI\xfa\xf4O\xcc\xec\xd3\xf7\x92\xf4'ox\x8b:=~\x9f$\xc9/\x8f\xff:K\xd3_|\xf1e5`\xd9\xf0\xd6s~\x9f\xa4\xc7\xef\x92\xe4\xc3\xe3?\x9f\xa4\xe9O\x03\xfd\xbb$=K\x9f\xfe\x91uH\x82\x9eZb\x9e$\xc9\x0f\xef\xff\x8a4I\u04f34\xcd\xccM\x8f\xba\xac\x06L\u007fx\xffg=\xab\xbeb;\x84\xed\xc0\xdb:\xcf/.\xe0w@\xf7C%\x95B\xddKQk0\x12\x18\xfcA\xbaC%\xc1e\xfe\x9c\xfem\xe0\xdb<\xa3\xeb\x05\xebp\x03\xfeO\x1b\xc5\xc5.\xcfPT\xb2\xe6b7\x01\xcf_{J\x9eqaP\xf5\n\r3\\\x8a\x97\x1bx\xfe6\xa2\xe4Y#U\xf7rb%\xee7Ruyf\xd8N\xbf\xb4\x17g_\xb8\x9b\xbe\xdcLW\x1e\xf2\x835\xe2\x156lh\rp\rf\x8f@*\u00a0\xb1\x86F*\u0426\xe6\x02\x98\xa8\xe9\x93\x1cL\t\x9f\xef\x114\x1a\xc3\xc5NC\x8d=\x8a\x9a\xa4H\x11\xb8;Y\x93\xd5^\xf0\x06\xac\xfd\xf0A\xec\x80\xf3\xe27\x05\u070e\xda\x1cf\xfe|+\x1a\t56\\\xa0\x86\xbd\xfc\x06\x98\x13\xcb5X7am\x15\x9a\u0702\xb5w11Zk\xed\xb7<\xab\x99a\xc1+\xe7F\r\b\xb7\u0430Vc\x9e)lP\xa1\xa8PoN\xc1\xea\xa6j\x1d\xb0\xc2iU\xe3\xe4y:\xb1\x95\xb2\xcd3\xd9\xd3w\xd6:\x16G\xab\xa4\xd0F1.L8\xf7\x15b\xef\xfd\xa27\x9e\xc6E%\xbb\xbeEc\xc3\xc2\u04fa^*3j\xe0h\xda(d\u0768\x94\xa3\u0572\xd2\xc1DGc\xc6(\xbe\x1d\x8c3\xc0\u049c{\xe9]4=\x1e=\x9c\xd3\xc1>r\xcd\x1b\xeb\v\x03\xb2G\u015c%\xeet\x99_\\\x10\xeb\xe7{\xd4\b\x06\xbb\xbee\x0650\x85\xf6\x01\x04\xbd\x86\x91\xb0E\x18\x04o8\u04bb\x0036\x18\x94\x94\x06d\x03f\xcf5\t\xa9\xa4h\xf8np7\x94\xb9\xbd\xc0\xbe\x17\x17\xfd`\\\x9c\xb6h\xe0\x1a.\xed\xe7\u023a\xc5#d\x91\x99K\xf0\x90gY\x88?++d\xd8yQ\rH\xb1wE\xf4\xb2,G\x86\x10C\xd7y`\xd0^@5P\xd4R\xaa\xe9RW{\xec\x98\x17A\xbcxmPh\x17\x12\xf6tQ\xfeUKQ\xf8o\x8b\x1c&\x1d\xd8`\xe4\xa4\xc4\xc1\xb1\u0730\xae},\xcb\xe38\x0e\x94\xf7\x19^St\xcd\x1c~\xf5\u045a\u02fdS\xcfW]\xbe\x04\x1fp\xb9\xf5\xc6\xfd>\xbf\xfa\xe8\x01\xafS>\a\x9f\x1f\xf2L\x0e\xbd\x89\x02\xe7\xea\xe3\x1f\u01ce\xb9V\x1f?V+\xfc\x9a\xea@\xd0\xe9\x93\xff\xb5o\x1f\x0e\xe7\xabO\x1e0\xa2\xe1\x94\xf2s+jl\xe6F|\xfa\xff\xcf\u026bO\x1f\x99\x95c\x87{=&'t\xac\u05ee\x99\x84\x84\xa5\xf2\xe5\u02e1\x83zEe\xd0p\xaa~\x8b\xbc.\x8ay\x97\xbd\u02b3\x82\x86\x83\x89H\xfd\x96\byH\xff@'\xc2\b\xb4\x1e\x99\x80\x96\x90\xb6\x0eL1\"\xeeD|\xc9\b\u0488\x90O\x85a\x050\xd7&\x06\f^\x1b\x02v2Xg\x81\x9d$r\xaf\xa4\x91s}-\xc1J\xc2k3\xa2\x93\xa4\x18\xdd\xcet\x0eh\x9eQK\xf9\xec\xd5g\x1b C4\xfe\xed\x85%\x15\xe5\xc801m\xb9\xe8\xb7pq\x01[.\x98\xba\xe9\xb7\u04e80\x0eH\xc0E\xcd+\u05d5\xdc\x03R40c[\x9b\xc2^\xa1FA\xe3\n0z\u069db]\x99O\xe3\xd5\x06\x9e]\x16\x85\x13) \x1e\xac\xa0F\x83\xaa\x9b\xcd!\x15*\u00f8\x18\xe5\x80\xde\u02e1\xad\xa9\xfbE\xd3\xc8\xc5\x05\xbc\x91\n\xc6\x11\xf6\x05\xd8\x1a\u0471\x9b\xc5I`\u0509u\xa5\xf8\xd6\xe9\xe7\"\xf8\x05|\xb3\xe7\xd5\x1e\xb8\xd1\xd86\xb6s2A\xac\x95\x14_\xa32\xae\xe52\xf8\xfd_^{\x8e2_\u0304\u04d8g'\xc1y\xd0zzcG\xd2hd\x1cG\xaf\u0160V4R\xbahv\x83\xa6\xe3*\xdc\u0145\u007f\x0ez+\x97]\x95\xec:\x1a\xcfZ.\u0411\x8d<\xcd+\x02lF91.\x99\x9d\xf4I2\xa5\xf0N\xb1~\x1f\xa1\x96\xe2\xc0\x9a\xed\"\xa8f\xbb\x110l\x81\x18/\xd0\u058bo\xf3y\xed\xb1\xa5\u01c2d\xe5\t\xeaM\xf7p\xbb\x8a\xb7\xee\x00\xa5\xd8\tn3\xd4\xc26\xf8Op\x97A\xf6\xc0\x94!'\x87B\xaa\xd1A\x9b,\xfd\x96&b;\xa9#7{T\xe4\xe81\x17|\xba\xc0(\xe2\x05\xc8\b\u03f3~\xbb\x81\xf3\xf8\x16\xf7W\x8c\x99V\xe4\xa7#EA\xf7\xc3-\xac1>\xbb\xbc\x9f\u0552\xbd\x95\xab\x06\x16\u04c3Y=\u00a39\xb1'<\x8e|'\xd7\xee\u010d\xde@\xda!\xee2.\x9b\"3\xcbZf\xaf\u0651\xd3}W$\xd6\x1fE\uae05y\xb94\xa89\xfc\x84\xdd\xcep+\x17F3\x95\x8f\xcey6\x9d\b\n\a\xdeE\x9c\xecQ\xb0\x9e\xdf!\u02e3\xef \xc8\xd5\a\u06e0\xa7\xa5\xce7j*\u042cm\x1dX\xc2[\x03\xb5D\rB\x1a\xe0\xa2j\x87\x1a\xddN)U\ao_\x95\xb9=g\x15\xb2\x1b-\xed\xee\x97\xd3Z;\xd5/\xab=5\uaaf5\xea\x02\x93\x96\xde\x15p\v\x85\x9d~\uc9f1\xba,\x96\xad\xe5@\x16\xafl\xcbI'^\x10\x97h\xbc*~\x18\xc1\xbf\x86\x0f\x96\x94<[,\x92Ky\xf1J\xb9D\xe3Er\x81\x1e\xa8\u038bqZ\x9d\x0fQ'\xfe\xf2>:\xb9o\u076a \xff\xa4\x80\x87\ap\xbe&\xafS\xe1v\xffm\xee.\x16w\xd2\xf9\xc4\xe7\ubfbeW\x9b\x85\x1f\xd7\xfd\xb7\xee\xb7`O\xd4stim\x98\xd9\xf6\xec2\x84\xd0\xf8#\u009cy\u0797hu\xd8-\xfd\xf2\xec\u04b7\xb1X\xdbQ\xad\xe8W\x8b\u026e\xf9\xaf\x15\xab\x06\xac\xfae\xd2\xeb\x90\xc7S\xf5\xd4#\xc7$\b\x16\x84\x0e\x19\x96\x9fE\xb6\xb8$\x81\xdb\xf1\xdd\xe6\v\u00e8\xc7|O\b\xc2C\xfb\x8c\x9d\x1b\xa9Ai\xe8$\xc7\x1dyU\x9f\xe9`\xe89\xab\xe7\x82\x0e\xf3V\xf3\xc0Q#\xbb\xfb\xee\x0e\ag-}\x91c\xef4\x06D\xd2\xef\x98\tf/\xb0\xb4\x85\x1a\xfd}b\xe6={MJhy\v\xe5O\f=\xe4q\x9fxD\xad\xb6\u02d4k\x82\xf1-\u02eev\xa7\x03\xef\xed_\xef\u0335\xea\xace4\x1d\xf2$\xf9o\x00\x00\x00\xff\xff\x85i\xbeS\xb4\x16\x00\x00")
