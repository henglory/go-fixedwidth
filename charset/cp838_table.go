package charset

var ebc838table = [][]byte{
	[]byte{'\x00'},
	[]byte{'\x01'},
	[]byte{'\x02'},
	[]byte{'\x03'},
	[]byte{'\xC2', '\x9C'},
	[]byte{'\x09'},
	[]byte{'\xC2', '\x86'},
	[]byte{'\x7F'},
	[]byte{'\xC2', '\x97'},
	[]byte{'\xC2', '\x8D'},
	[]byte{'\xC2', '\x8E'},
	[]byte{'\x0B'},
	[]byte{'\x0C'},
	[]byte{'\x0D'},
	[]byte{'\x0E'},
	[]byte{'\x0F'},
	[]byte{'\x10'},
	[]byte{'\x11'},
	[]byte{'\x12'},
	[]byte{'\x13'},
	[]byte{'\xC2', '\x9D'},
	[]byte{'\x0A'},
	[]byte{'\x08'},
	[]byte{'\xC2', '\x87'},
	[]byte{'\x18'},
	[]byte{'\x19'},
	[]byte{'\xC2', '\x92'},
	[]byte{'\xC2', '\x8F'},
	[]byte{'\x1C'},
	[]byte{'\x1D'},
	[]byte{'\x1E'},
	[]byte{'\x1F'},
	[]byte{'\xC2', '\x80'},
	[]byte{'\xC2', '\x81'},
	[]byte{'\xC2', '\x82'},
	[]byte{'\xC2', '\x83'},
	[]byte{'\xC2', '\x84'},
	[]byte{'\x0A'},
	[]byte{'\x17'},
	[]byte{'\x1B'},
	[]byte{'\xC2', '\x88'},
	[]byte{'\xC2', '\x89'},
	[]byte{'\xC2', '\x8A'},
	[]byte{'\xC2', '\x8B'},
	[]byte{'\xC2', '\x8C'},
	[]byte{'\x05'},
	[]byte{'\x06'},
	[]byte{'\x07'},
	[]byte{'\xC2', '\x90'},
	[]byte{'\xC2', '\x91'},
	[]byte{'\x16'},
	[]byte{'\xC2', '\x93'},
	[]byte{'\xC2', '\x94'},
	[]byte{'\xC2', '\x95'},
	[]byte{'\xC2', '\x96'},
	[]byte{'\x04'},
	[]byte{'\xC2', '\x98'},
	[]byte{'\xC2', '\x99'},
	[]byte{'\xC2', '\x9A'},
	[]byte{'\xC2', '\x9B'},
	[]byte{'\x14'},
	[]byte{'\x15'},
	[]byte{'\xC2', '\x9E'},
	[]byte{'\x1A'},
	[]byte{'\x20'},
	[]byte{'\xC2', '\xA0'},
	[]byte{'\xE0', '\xB8', '\x81'},
	[]byte{'\xE0', '\xB8', '\x82'},
	[]byte{'\xE0', '\xB8', '\x83'},
	[]byte{'\xE0', '\xB8', '\x84'},
	[]byte{'\xE0', '\xB8', '\x85'},
	[]byte{'\xE0', '\xB8', '\x86'},
	[]byte{'\xE0', '\xB8', '\x87'},
	[]byte{'\x5B'},
	[]byte{'\xC2', '\xA2'},
	[]byte{'\x2E'},
	[]byte{'\x3C'},
	[]byte{'\x28'},
	[]byte{'\x2B'},
	[]byte{'\x7C'},
	[]byte{'\x26'},
	[]byte{'\xE0', '\xB9', '\x88'},
	[]byte{'\xE0', '\xB8', '\x88'},
	[]byte{'\xE0', '\xB8', '\x89'},
	[]byte{'\xE0', '\xB8', '\x8A'},
	[]byte{'\xE0', '\xB8', '\x8B'},
	[]byte{'\xE0', '\xB8', '\x8C'},
	[]byte{'\xE0', '\xB8', '\x8D'},
	[]byte{'\xE0', '\xB8', '\x8E'},
	[]byte{'\x5D'},
	[]byte{'\x21'},
	[]byte{'\x24'},
	[]byte{'\x2A'},
	[]byte{'\x29'},
	[]byte{'\x3B'},
	[]byte{'\xC2', '\xAC'},
	[]byte{'\x2D'},
	[]byte{'\x2F'},
	[]byte{'\xE0', '\xB8', '\x8F'},
	[]byte{'\xE0', '\xB8', '\x90'},
	[]byte{'\xE0', '\xB8', '\x91'},
	[]byte{'\xE0', '\xB8', '\x92'},
	[]byte{'\xE0', '\xB8', '\x93'},
	[]byte{'\xE0', '\xB8', '\x94'},
	[]byte{'\xE0', '\xB8', '\x95'},
	[]byte{'\x5E'},
	[]byte{'\xC2', '\xA6'},
	[]byte{'\x2C'},
	[]byte{'\x25'},
	[]byte{'\x5F'},
	[]byte{'\x3E'},
	[]byte{'\x3F'},
	[]byte{'\xE0', '\xB8', '\xBF'},
	[]byte{'\xE0', '\xB9', '\x8E'},
	[]byte{'\xE0', '\xB8', '\x96'},
	[]byte{'\xE0', '\xB8', '\x97'},
	[]byte{'\xE0', '\xB8', '\x98'},
	[]byte{'\xE0', '\xB8', '\x99'},
	[]byte{'\xE0', '\xB8', '\x9A'},
	[]byte{'\xE0', '\xB8', '\x9B'},
	[]byte{'\xE0', '\xB8', '\x9C'},
	[]byte{'\x60'},
	[]byte{'\x3A'},
	[]byte{'\x23'},
	[]byte{'\x40'},
	[]byte{'\x27'},
	[]byte{'\x3D'},
	[]byte{'\x22'},
	[]byte{'\xE0', '\xB9', '\x8F'},
	[]byte{'\x61'},
	[]byte{'\x62'},
	[]byte{'\x63'},
	[]byte{'\x64'},
	[]byte{'\x65'},
	[]byte{'\x66'},
	[]byte{'\x67'},
	[]byte{'\x68'},
	[]byte{'\x69'},
	[]byte{'\xE0', '\xB8', '\x9D'},
	[]byte{'\xE0', '\xB8', '\x9E'},
	[]byte{'\xE0', '\xB8', '\x9F'},
	[]byte{'\xE0', '\xB8', '\xA0'},
	[]byte{'\xE0', '\xB8', '\xA1'},
	[]byte{'\xE0', '\xB8', '\xA2'},
	[]byte{'\xE0', '\xB9', '\x9A'},
	[]byte{'\x6A'},
	[]byte{'\x6B'},
	[]byte{'\x6C'},
	[]byte{'\x6D'},
	[]byte{'\x6E'},
	[]byte{'\x6F'},
	[]byte{'\x70'},
	[]byte{'\x71'},
	[]byte{'\x72'},
	[]byte{'\xE0', '\xB8', '\xA3'},
	[]byte{'\xE0', '\xB8', '\xA4'},
	[]byte{'\xE0', '\xB8', '\xA5'},
	[]byte{'\xE0', '\xB8', '\xA6'},
	[]byte{'\xE0', '\xB8', '\xA7'},
	[]byte{'\xE0', '\xB8', '\xA8'},
	[]byte{'\xE0', '\xB9', '\x9B'},
	[]byte{'\x7E'},
	[]byte{'\x73'},
	[]byte{'\x74'},
	[]byte{'\x75'},
	[]byte{'\x76'},
	[]byte{'\x77'},
	[]byte{'\x78'},
	[]byte{'\x79'},
	[]byte{'\x7A'},
	[]byte{'\xE0', '\xB8', '\xA9'},
	[]byte{'\xE0', '\xB8', '\xAA'},
	[]byte{'\xE0', '\xB8', '\xAB'},
	[]byte{'\xE0', '\xB8', '\xAC'},
	[]byte{'\xE0', '\xB8', '\xAD'},
	[]byte{'\xE0', '\xB8', '\xAE'},
	[]byte{'\xE0', '\xB9', '\x90'},
	[]byte{'\xE0', '\xB9', '\x91'},
	[]byte{'\xE0', '\xB9', '\x92'},
	[]byte{'\xE0', '\xB9', '\x93'},
	[]byte{'\xE0', '\xB9', '\x94'},
	[]byte{'\xE0', '\xB9', '\x95'},
	[]byte{'\xE0', '\xB9', '\x96'},
	[]byte{'\xE0', '\xB9', '\x97'},
	[]byte{'\xE0', '\xB9', '\x98'},
	[]byte{'\xE0', '\xB9', '\x99'},
	[]byte{'\xE0', '\xB8', '\xAF'},
	[]byte{'\xE0', '\xB8', '\xB0'},
	[]byte{'\xE0', '\xB8', '\xB1'},
	[]byte{'\xE0', '\xB8', '\xB2'},
	[]byte{'\xE0', '\xB8', '\xB3'},
	[]byte{'\xE0', '\xB8', '\xB4'},
	[]byte{'\x7B'},
	[]byte{'\x41'},
	[]byte{'\x42'},
	[]byte{'\x43'},
	[]byte{'\x44'},
	[]byte{'\x45'},
	[]byte{'\x46'},
	[]byte{'\x47'},
	[]byte{'\x48'},
	[]byte{'\x49'},
	[]byte{'\xE0', '\xB9', '\x89'},
	[]byte{'\xE0', '\xB8', '\xB5'},
	[]byte{'\xE0', '\xB8', '\xB6'},
	[]byte{'\xE0', '\xB8', '\xB7'},
	[]byte{'\xE0', '\xB8', '\xB8'},
	[]byte{'\xE0', '\xB8', '\xB9'},
	[]byte{'\x7D'},
	[]byte{'\x4A'},
	[]byte{'\x4B'},
	[]byte{'\x4C'},
	[]byte{'\x4D'},
	[]byte{'\x4E'},
	[]byte{'\x4F'},
	[]byte{'\x50'},
	[]byte{'\x51'},
	[]byte{'\x52'},
	[]byte{'\xE0', '\xB8', '\xBA'},
	[]byte{'\xE0', '\xB9', '\x80'},
	[]byte{'\xE0', '\xB9', '\x81'},
	[]byte{'\xE0', '\xB9', '\x82'},
	[]byte{'\xE0', '\xB9', '\x83'},
	[]byte{'\xE0', '\xB9', '\x84'},
	[]byte{'\x5C'},
	[]byte{'\xE0', '\xB9', '\x8A'},
	[]byte{'\x53'},
	[]byte{'\x54'},
	[]byte{'\x55'},
	[]byte{'\x56'},
	[]byte{'\x57'},
	[]byte{'\x58'},
	[]byte{'\x59'},
	[]byte{'\x5A'},
	[]byte{'\xE0', '\xB9', '\x85'},
	[]byte{'\xE0', '\xB9', '\x86'},
	[]byte{'\xE0', '\xB9', '\x87'},
	[]byte{'\xE0', '\xB9', '\x88'},
	[]byte{'\xE0', '\xB9', '\x89'},
	[]byte{'\xE0', '\xB9', '\x8A'},
	[]byte{'\x30'},
	[]byte{'\x31'},
	[]byte{'\x32'},
	[]byte{'\x33'},
	[]byte{'\x34'},
	[]byte{'\x35'},
	[]byte{'\x36'},
	[]byte{'\x37'},
	[]byte{'\x38'},
	[]byte{'\x39'},
	[]byte{'\xE0', '\xB9', '\x8B'},
	[]byte{'\xE0', '\xB9', '\x8C'},
	[]byte{'\xE0', '\xB9', '\x8D'},
	[]byte{'\xE0', '\xB9', '\x8B'},
	[]byte{'\xE0', '\xB9', '\x8C'},
	[]byte{'\xC2', '\x9F'},
}

var utf8Toebc838Map = map[string]byte{
	string([]byte{'\x00'}):                 '\x00',
	string([]byte{'\x01'}):                 '\x01',
	string([]byte{'\x02'}):                 '\x02',
	string([]byte{'\x03'}):                 '\x03',
	string([]byte{'\xC2', '\x9C'}):         '\x04',
	string([]byte{'\x09'}):                 '\x05',
	string([]byte{'\xC2', '\x86'}):         '\x06',
	string([]byte{'\x7F'}):                 '\x07',
	string([]byte{'\xC2', '\x97'}):         '\x08',
	string([]byte{'\xC2', '\x8D'}):         '\x09',
	string([]byte{'\xC2', '\x8E'}):         '\x0A',
	string([]byte{'\x0B'}):                 '\x0B',
	string([]byte{'\x0C'}):                 '\x0C',
	string([]byte{'\x0D'}):                 '\x0D',
	string([]byte{'\x0E'}):                 '\x0E',
	string([]byte{'\x0F'}):                 '\x0F',
	string([]byte{'\x10'}):                 '\x10',
	string([]byte{'\x11'}):                 '\x11',
	string([]byte{'\x12'}):                 '\x12',
	string([]byte{'\x13'}):                 '\x13',
	string([]byte{'\xC2', '\x9D'}):         '\x14',
	string([]byte{'\x0A'}):                 '\x15',
	string([]byte{'\x08'}):                 '\x16',
	string([]byte{'\xC2', '\x87'}):         '\x17',
	string([]byte{'\x18'}):                 '\x18',
	string([]byte{'\x19'}):                 '\x19',
	string([]byte{'\xC2', '\x92'}):         '\x1A',
	string([]byte{'\xC2', '\x8F'}):         '\x1B',
	string([]byte{'\x1C'}):                 '\x1C',
	string([]byte{'\x1D'}):                 '\x1D',
	string([]byte{'\x1E'}):                 '\x1E',
	string([]byte{'\x1F'}):                 '\x1F',
	string([]byte{'\xC2', '\x80'}):         '\x20',
	string([]byte{'\xC2', '\x81'}):         '\x21',
	string([]byte{'\xC2', '\x82'}):         '\x22',
	string([]byte{'\xC2', '\x83'}):         '\x23',
	string([]byte{'\xC2', '\x84'}):         '\x24',
	string([]byte{'\x0A'}):                 '\x25',
	string([]byte{'\x17'}):                 '\x26',
	string([]byte{'\x1B'}):                 '\x27',
	string([]byte{'\xC2', '\x88'}):         '\x28',
	string([]byte{'\xC2', '\x89'}):         '\x29',
	string([]byte{'\xC2', '\x8A'}):         '\x2A',
	string([]byte{'\xC2', '\x8B'}):         '\x2B',
	string([]byte{'\xC2', '\x8C'}):         '\x2C',
	string([]byte{'\x05'}):                 '\x2D',
	string([]byte{'\x06'}):                 '\x2E',
	string([]byte{'\x07'}):                 '\x2F',
	string([]byte{'\xC2', '\x90'}):         '\x30',
	string([]byte{'\xC2', '\x91'}):         '\x31',
	string([]byte{'\x16'}):                 '\x32',
	string([]byte{'\xC2', '\x93'}):         '\x33',
	string([]byte{'\xC2', '\x94'}):         '\x34',
	string([]byte{'\xC2', '\x95'}):         '\x35',
	string([]byte{'\xC2', '\x96'}):         '\x36',
	string([]byte{'\x04'}):                 '\x37',
	string([]byte{'\xC2', '\x98'}):         '\x38',
	string([]byte{'\xC2', '\x99'}):         '\x39',
	string([]byte{'\xC2', '\x9A'}):         '\x3A',
	string([]byte{'\xC2', '\x9B'}):         '\x3B',
	string([]byte{'\x14'}):                 '\x3C',
	string([]byte{'\x15'}):                 '\x3D',
	string([]byte{'\xC2', '\x9E'}):         '\x3E',
	string([]byte{'\x1A'}):                 '\x3F',
	string([]byte{'\x20'}):                 '\x40',
	string([]byte{'\xC2', '\xA0'}):         '\x41',
	string([]byte{'\xE0', '\xB8', '\x81'}): '\x42',
	string([]byte{'\xE0', '\xB8', '\x82'}): '\x43',
	string([]byte{'\xE0', '\xB8', '\x83'}): '\x44',
	string([]byte{'\xE0', '\xB8', '\x84'}): '\x45',
	string([]byte{'\xE0', '\xB8', '\x85'}): '\x46',
	string([]byte{'\xE0', '\xB8', '\x86'}): '\x47',
	string([]byte{'\xE0', '\xB8', '\x87'}): '\x48',
	string([]byte{'\x5B'}):                 '\x49',
	string([]byte{'\xC2', '\xA2'}):         '\x4A',
	string([]byte{'\x2E'}):                 '\x4B',
	string([]byte{'\x3C'}):                 '\x4C',
	string([]byte{'\x28'}):                 '\x4D',
	string([]byte{'\x2B'}):                 '\x4E',
	string([]byte{'\x7C'}):                 '\x4F',
	string([]byte{'\x26'}):                 '\x50',
	string([]byte{'\xE0', '\xB9', '\x88'}): '\x51',
	string([]byte{'\xE0', '\xB8', '\x88'}): '\x52',
	string([]byte{'\xE0', '\xB8', '\x89'}): '\x53',
	string([]byte{'\xE0', '\xB8', '\x8A'}): '\x54',
	string([]byte{'\xE0', '\xB8', '\x8B'}): '\x55',
	string([]byte{'\xE0', '\xB8', '\x8C'}): '\x56',
	string([]byte{'\xE0', '\xB8', '\x8D'}): '\x57',
	string([]byte{'\xE0', '\xB8', '\x8E'}): '\x58',
	string([]byte{'\x5D'}):                 '\x59',
	string([]byte{'\x21'}):                 '\x5A',
	string([]byte{'\x24'}):                 '\x5B',
	string([]byte{'\x2A'}):                 '\x5C',
	string([]byte{'\x29'}):                 '\x5D',
	string([]byte{'\x3B'}):                 '\x5E',
	string([]byte{'\xC2', '\xAC'}):         '\x5F',
	string([]byte{'\x2D'}):                 '\x60',
	string([]byte{'\x2F'}):                 '\x61',
	string([]byte{'\xE0', '\xB8', '\x8F'}): '\x62',
	string([]byte{'\xE0', '\xB8', '\x90'}): '\x63',
	string([]byte{'\xE0', '\xB8', '\x91'}): '\x64',
	string([]byte{'\xE0', '\xB8', '\x92'}): '\x65',
	string([]byte{'\xE0', '\xB8', '\x93'}): '\x66',
	string([]byte{'\xE0', '\xB8', '\x94'}): '\x67',
	string([]byte{'\xE0', '\xB8', '\x95'}): '\x68',
	string([]byte{'\x5E'}):                 '\x69',
	string([]byte{'\xC2', '\xA6'}):         '\x6A',
	string([]byte{'\x2C'}):                 '\x6B',
	string([]byte{'\x25'}):                 '\x6C',
	string([]byte{'\x5F'}):                 '\x6D',
	string([]byte{'\x3E'}):                 '\x6E',
	string([]byte{'\x3F'}):                 '\x6F',
	string([]byte{'\xE0', '\xB8', '\xBF'}): '\x70',
	string([]byte{'\xE0', '\xB9', '\x8E'}): '\x71',
	string([]byte{'\xE0', '\xB8', '\x96'}): '\x72',
	string([]byte{'\xE0', '\xB8', '\x97'}): '\x73',
	string([]byte{'\xE0', '\xB8', '\x98'}): '\x74',
	string([]byte{'\xE0', '\xB8', '\x99'}): '\x75',
	string([]byte{'\xE0', '\xB8', '\x9A'}): '\x76',
	string([]byte{'\xE0', '\xB8', '\x9B'}): '\x77',
	string([]byte{'\xE0', '\xB8', '\x9C'}): '\x78',
	string([]byte{'\x60'}):                 '\x79',
	string([]byte{'\x3A'}):                 '\x7A',
	string([]byte{'\x23'}):                 '\x7B',
	string([]byte{'\x40'}):                 '\x7C',
	string([]byte{'\x27'}):                 '\x7D',
	string([]byte{'\x3D'}):                 '\x7E',
	string([]byte{'\x22'}):                 '\x7F',
	string([]byte{'\xE0', '\xB9', '\x8F'}): '\x80',
	string([]byte{'\x61'}):                 '\x81',
	string([]byte{'\x62'}):                 '\x82',
	string([]byte{'\x63'}):                 '\x83',
	string([]byte{'\x64'}):                 '\x84',
	string([]byte{'\x65'}):                 '\x85',
	string([]byte{'\x66'}):                 '\x86',
	string([]byte{'\x67'}):                 '\x87',
	string([]byte{'\x68'}):                 '\x88',
	string([]byte{'\x69'}):                 '\x89',
	string([]byte{'\xE0', '\xB8', '\x9D'}): '\x8A',
	string([]byte{'\xE0', '\xB8', '\x9E'}): '\x8B',
	string([]byte{'\xE0', '\xB8', '\x9F'}): '\x8C',
	string([]byte{'\xE0', '\xB8', '\xA0'}): '\x8D',
	string([]byte{'\xE0', '\xB8', '\xA1'}): '\x8E',
	string([]byte{'\xE0', '\xB8', '\xA2'}): '\x8F',
	string([]byte{'\xE0', '\xB9', '\x9A'}): '\x90',
	string([]byte{'\x6A'}):                 '\x91',
	string([]byte{'\x6B'}):                 '\x92',
	string([]byte{'\x6C'}):                 '\x93',
	string([]byte{'\x6D'}):                 '\x94',
	string([]byte{'\x6E'}):                 '\x95',
	string([]byte{'\x6F'}):                 '\x96',
	string([]byte{'\x70'}):                 '\x97',
	string([]byte{'\x71'}):                 '\x98',
	string([]byte{'\x72'}):                 '\x99',
	string([]byte{'\xE0', '\xB8', '\xA3'}): '\x9A',
	string([]byte{'\xE0', '\xB8', '\xA4'}): '\x9B',
	string([]byte{'\xE0', '\xB8', '\xA5'}): '\x9C',
	string([]byte{'\xE0', '\xB8', '\xA6'}): '\x9D',
	string([]byte{'\xE0', '\xB8', '\xA7'}): '\x9E',
	string([]byte{'\xE0', '\xB8', '\xA8'}): '\x9F',
	string([]byte{'\xE0', '\xB9', '\x9B'}): '\xA0',
	string([]byte{'\x7E'}):                 '\xA1',
	string([]byte{'\x73'}):                 '\xA2',
	string([]byte{'\x74'}):                 '\xA3',
	string([]byte{'\x75'}):                 '\xA4',
	string([]byte{'\x76'}):                 '\xA5',
	string([]byte{'\x77'}):                 '\xA6',
	string([]byte{'\x78'}):                 '\xA7',
	string([]byte{'\x79'}):                 '\xA8',
	string([]byte{'\x7A'}):                 '\xA9',
	string([]byte{'\xE0', '\xB8', '\xA9'}): '\xAA',
	string([]byte{'\xE0', '\xB8', '\xAA'}): '\xAB',
	string([]byte{'\xE0', '\xB8', '\xAB'}): '\xAC',
	string([]byte{'\xE0', '\xB8', '\xAC'}): '\xAD',
	string([]byte{'\xE0', '\xB8', '\xAD'}): '\xAE',
	string([]byte{'\xE0', '\xB8', '\xAE'}): '\xAF',
	string([]byte{'\xE0', '\xB9', '\x90'}): '\xB0',
	string([]byte{'\xE0', '\xB9', '\x91'}): '\xB1',
	string([]byte{'\xE0', '\xB9', '\x92'}): '\xB2',
	string([]byte{'\xE0', '\xB9', '\x93'}): '\xB3',
	string([]byte{'\xE0', '\xB9', '\x94'}): '\xB4',
	string([]byte{'\xE0', '\xB9', '\x95'}): '\xB5',
	string([]byte{'\xE0', '\xB9', '\x96'}): '\xB6',
	string([]byte{'\xE0', '\xB9', '\x97'}): '\xB7',
	string([]byte{'\xE0', '\xB9', '\x98'}): '\xB8',
	string([]byte{'\xE0', '\xB9', '\x99'}): '\xB9',
	string([]byte{'\xE0', '\xB8', '\xAF'}): '\xBA',
	string([]byte{'\xE0', '\xB8', '\xB0'}): '\xBB',
	string([]byte{'\xE0', '\xB8', '\xB1'}): '\xBC',
	string([]byte{'\xE0', '\xB8', '\xB2'}): '\xBD',
	string([]byte{'\xE0', '\xB8', '\xB3'}): '\xBE',
	string([]byte{'\xE0', '\xB8', '\xB4'}): '\xBF',
	string([]byte{'\x7B'}):                 '\xC0',
	string([]byte{'\x41'}):                 '\xC1',
	string([]byte{'\x42'}):                 '\xC2',
	string([]byte{'\x43'}):                 '\xC3',
	string([]byte{'\x44'}):                 '\xC4',
	string([]byte{'\x45'}):                 '\xC5',
	string([]byte{'\x46'}):                 '\xC6',
	string([]byte{'\x47'}):                 '\xC7',
	string([]byte{'\x48'}):                 '\xC8',
	string([]byte{'\x49'}):                 '\xC9',
	string([]byte{'\xE0', '\xB9', '\x89'}): '\xCA',
	string([]byte{'\xE0', '\xB8', '\xB5'}): '\xCB',
	string([]byte{'\xE0', '\xB8', '\xB6'}): '\xCC',
	string([]byte{'\xE0', '\xB8', '\xB7'}): '\xCD',
	string([]byte{'\xE0', '\xB8', '\xB8'}): '\xCE',
	string([]byte{'\xE0', '\xB8', '\xB9'}): '\xCF',
	string([]byte{'\x7D'}):                 '\xD0',
	string([]byte{'\x4A'}):                 '\xD1',
	string([]byte{'\x4B'}):                 '\xD2',
	string([]byte{'\x4C'}):                 '\xD3',
	string([]byte{'\x4D'}):                 '\xD4',
	string([]byte{'\x4E'}):                 '\xD5',
	string([]byte{'\x4F'}):                 '\xD6',
	string([]byte{'\x50'}):                 '\xD7',
	string([]byte{'\x51'}):                 '\xD8',
	string([]byte{'\x52'}):                 '\xD9',
	string([]byte{'\xE0', '\xB8', '\xBA'}): '\xDA',
	string([]byte{'\xE0', '\xB9', '\x80'}): '\xDB',
	string([]byte{'\xE0', '\xB9', '\x81'}): '\xDC',
	string([]byte{'\xE0', '\xB9', '\x82'}): '\xDD',
	string([]byte{'\xE0', '\xB9', '\x83'}): '\xDE',
	string([]byte{'\xE0', '\xB9', '\x84'}): '\xDF',
	string([]byte{'\x5C'}):                 '\xE0',
	string([]byte{'\xE0', '\xB9', '\x8A'}): '\xE1',
	string([]byte{'\x53'}):                 '\xE2',
	string([]byte{'\x54'}):                 '\xE3',
	string([]byte{'\x55'}):                 '\xE4',
	string([]byte{'\x56'}):                 '\xE5',
	string([]byte{'\x57'}):                 '\xE6',
	string([]byte{'\x58'}):                 '\xE7',
	string([]byte{'\x59'}):                 '\xE8',
	string([]byte{'\x5A'}):                 '\xE9',
	string([]byte{'\xE0', '\xB9', '\x85'}): '\xEA',
	string([]byte{'\xE0', '\xB9', '\x86'}): '\xEB',
	string([]byte{'\xE0', '\xB9', '\x87'}): '\xEC',
	string([]byte{'\xE0', '\xB9', '\x88'}): '\xED',
	string([]byte{'\xE0', '\xB9', '\x89'}): '\xEE',
	string([]byte{'\xE0', '\xB9', '\x8A'}): '\xEF',
	string([]byte{'\x30'}):                 '\xF0',
	string([]byte{'\x31'}):                 '\xF1',
	string([]byte{'\x32'}):                 '\xF2',
	string([]byte{'\x33'}):                 '\xF3',
	string([]byte{'\x34'}):                 '\xF4',
	string([]byte{'\x35'}):                 '\xF5',
	string([]byte{'\x36'}):                 '\xF6',
	string([]byte{'\x37'}):                 '\xF7',
	string([]byte{'\x38'}):                 '\xF8',
	string([]byte{'\x39'}):                 '\xF9',
	string([]byte{'\xE0', '\xB9', '\x8B'}): '\xFA',
	string([]byte{'\xE0', '\xB9', '\x8C'}): '\xFB',
	string([]byte{'\xE0', '\xB9', '\x8D'}): '\xFC',
	string([]byte{'\xE0', '\xB9', '\x8B'}): '\xFD',
	string([]byte{'\xE0', '\xB9', '\x8C'}): '\xFE',
	string([]byte{'\xC2', '\x9F'}):         '\xFF',
}
