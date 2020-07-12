import ("github.com/jackpal/bencode-go")

type bencodeInfo struct {
	Peices string `bencode:"pieces"`
	PieceLength int `bencode:"piece length"`
	Length int `bencode:"length"`
	Name string `bencode:"name"`
}

type bencodeTorrent struct {
	Announce string `bencode:"anounce"`
	Info bencodeInfo `bencode:"info"`
}

func Open(r io.Reader) (*bencodeTorrent, error) {
	bto := bencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)

	if err != nil {
		return nil, err
	}

	return &bto, nil
}