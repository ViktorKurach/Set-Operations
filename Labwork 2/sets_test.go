package sets

import "testing"

type setAndSet struct {
	set1     Set
	set2     Set
	interRes Set
	unionRes Set
	difRes   Set
}

var testCases = []setAndSet{
	{
		CreateSetOfStrings([]string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing",
			"elit", "Aenean", "vestibulum", "eros", "magna", "vel", "accumsan", "sem", "faucibus", "et",
			"Morbi", "posuere", "enim", "in", "pulvinar", "dignissim", "Mauris", "velit", "orci", "facilisis",
			"in", "velit", "id", "vestibulum", "iaculis", "libero", "Integer", "orci", "eros", "tincidunt",
			"sollicitudin", "nulla", "ac", "porttitor", "ultricies", "metus", "Vivamus", "a", "nibh", "leo",
			"Donec", "dictum", "volutpat", "metus", "sed", "sagittis", "magna", "mattis", "placerat", "Nulla",
			"ac", "sollicitudin", "velit", "Duis", "sem", "augue", "aliquet", "eget", "neque", "at", "laoreet",
			"pulvinar", "orci", "Aenean", "nec", "mattis", "sapien", "Pellentesque", "tempus", "malesuada",
			"neque", "ut", "egestas", "tortor", "porttitor", "vel", "Donec", "in", "gravida", "augue", "eget",
			"luctus", "erat", "Duis", "laoreet", "porttitor", "pharetra", "Phasellus", "nec", "hendrerit",
			"urna", "Interdum", "et", "malesuada", "fames", "ac", "ante", "ipsum", "primis", "in", "faucibus",
			"Quisque", "ultricies", "vestibulum", "metus", "eu", "vulputate", "Integer", "accumsan", "aliquet",
			"gravida", "Morbi", "consectetur", "viverra", "nibh", "egestas", "vestibulum", "Nulla", "facilisi",
			"Duis", "tincidunt", "mi", "enim", "sit", "amet", "ultricies", "dui", "lobortis", "eu", "Nam",
			"dapibus", "nunc", "enim", "sed", "rhoncus", "magna", "rhoncus", "ac", "Phasellus", "tincidunt",
			"ante", "ac", "malesuada", "fringilla"}),
		CreateSetOfStrings([]string{"Aliquam", "ac", "erat", "tempus", "lacinia", "ex", "quis", "sagittis",
			"ipsum", "Nam", "neque", "augue", "rhoncus", "quis", "volutpat", "rhoncus", "sodales", "ut", "dolor",
			"Quisque", "at", "massa", "at", "diam", "suscipit", "posuere", "quis", "eget", "neque", "Phasellus",
			"eu", "euismod", "lacus", "Vivamus", "ac", "magna", "feugiat", "ullamcorper", "ex", "eget", "aliquam",
			"nisi", "Suspendisse", "aliquet", "nunc", "in", "elit", "facilisis", "ultrices", "Suspendisse", "non",
			"vehicula", "urna", "pulvinar", "fermentum", "diam", "Vivamus", "convallis", "leo", "sit", "amet",
			"turpis", "facilisis", "convallis", "Integer", "finibus", "enim", "odio", "Curabitur", "euismod",
			"nisi", "nulla", "vestibulum", "cursus", "diam", "ornare", "ut", "Nunc", "feugiat", "efficitur",
			"vehicula", "Class", "aptent", "taciti", "sociosqu", "ad", "litora", "torquent", "per", "conubia",
			"nostra", "per", "inceptos", "himenaeos", "Sed", "mattis", "ligula", "non", "velit", "ultricies",
			"lobortis", "Integer", "finibus", "tortor", "et", "velit", "accumsan", "congue", "Nam", "consectetur",
			"lacus", "efficitur", "sodales", "sollicitudin", "Phasellus", "non", "tempor", "velit", "Nullam", "id",
			"accumsan", "lacus", "Quisque", "est", "felis", "placerat", "in", "nunc", "et", "tempor", "scelerisque",
			"enim", "Quisque", "malesuada", "sodales", "lectus", "sit", "amet", "facilisis", "massa", "tincidunt",
			"a", "Suspendisse", "potenti", "Etiam", "vitae", "metus", "vehicula", "sagittis", "lacus", "a", "viverra",
			"purus", "Donec", "convallis", "eros", "mollis", "rhoncus", "magna", "id", "vulputate", "magna", "Donec",
			"consectetur", "in", "erat", "vitae", "sollicitudin", "Vestibulum", "facilisis", "et", "ligula", "sit",
			"amet", "tristique", "Ut", "dapibus", "finibus", "arcu", "ac", "rhoncus", "arcu", "placerat", "in",
			"Sed", "a", "venenatis", "augue", "Quisque", "quis", "sapien", "vel", "quam", "congue", "feugiat",
			"non", "eget", "mauris", "Vivamus", "ac", "elit", "eu", "nunc", "blandit", "tempus", "Nulla", "orci",
			"est", "blandit", "vehicula", "ipsum", "vitae", "molestie", "venenatis", "ante", "Ut", "ornare", "tempor",
			"est", "eget", "aliquet", "magna", "blandit", "vitae", "Phasellus", "in", "erat", "augue", "Praesent",
			"lacinia", "augue", "nibh", "ac", "feugiat", "nisi", "blandit", "nec", "Quisque", "nulla", "lacus",
			"molestie", "ut", "dignissim", "a", "pulvinar", "non", "sem", "Curabitur", "ut", "turpis", "pretium",
			"fringilla", "risus", "vitae", "tempor", "felis", "Phasellus", "at", "pharetra", "nisl", "Curabitur",
			"lobortis", "enim", "dui", "et", "cursus", "orci", "porta", "ut", "Nullam", "et", "vulputate", "neque",
			"Phasellus", "non", "convallis", "lectus", "Pellentesque", "semper", "cursus", "euismod", "Suspendisse",
			"aliquam", "aliquam", "magna", "ut", "euismod", "leo", "tincidunt", "ut"}),
		CreateSetOfStrings([]string{"ipsum", "dolor", "sit", "amet", "consectetur", "elit", "vestibulum", "eros",
			"magna", "vel", "accumsan", "sem", "et", "posuere", "enim", "in", "pulvinar", "dignissim", "velit", "orci",
			"facilisis", "id", "Integer", "tincidunt", "sollicitudin", "nulla", "ac", "ultricies", "metus", "Vivamus",
			"a", "nibh", "leo", "Donec", "volutpat", "sagittis", "mattis", "placerat", "Nulla", "augue", "aliquet",
			"eget", "neque", "at", "nec", "sapien", "Pellentesque", "tempus", "malesuada", "ut", "tortor", "erat",
			"pharetra", "Phasellus", "urna", "ante", "Quisque", "eu", "vulputate", "viverra", "dui", "lobortis", "Nam",
			"dapibus", "nunc", "rhoncus", "fringilla"}),
		CreateSetOfStrings([]string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing",
			"elit", "Aenean", "vestibulum", "eros", "magna", "vel", "accumsan", "sem", "faucibus", "et",
			"Morbi", "posuere", "enim", "in", "pulvinar", "dignissim", "Mauris", "velit", "orci", "facilisis",
			"in", "velit", "id", "vestibulum", "iaculis", "libero", "Integer", "orci", "eros", "tincidunt",
			"sollicitudin", "nulla", "ac", "porttitor", "ultricies", "metus", "Vivamus", "a", "nibh", "leo",
			"Donec", "dictum", "volutpat", "metus", "sed", "sagittis", "magna", "mattis", "placerat", "Nulla",
			"ac", "sollicitudin", "velit", "Duis", "sem", "augue", "aliquet", "eget", "neque", "at", "laoreet",
			"pulvinar", "orci", "Aenean", "nec", "mattis", "sapien", "Pellentesque", "tempus", "malesuada",
			"neque", "ut", "egestas", "tortor", "porttitor", "vel", "Donec", "in", "gravida", "augue", "eget",
			"luctus", "erat", "Duis", "laoreet", "porttitor", "pharetra", "Phasellus", "nec", "hendrerit",
			"urna", "Interdum", "et", "malesuada", "fames", "ac", "ante", "ipsum", "primis", "in", "faucibus",
			"Quisque", "ultricies", "vestibulum", "metus", "eu", "vulputate", "Integer", "accumsan", "aliquet",
			"gravida", "Morbi", "consectetur", "viverra", "nibh", "egestas", "vestibulum", "Nulla", "facilisi",
			"Duis", "tincidunt", "mi", "enim", "sit", "amet", "ultricies", "dui", "lobortis", "eu", "Nam",
			"dapibus", "nunc", "enim", "sed", "rhoncus", "magna", "rhoncus", "ac", "Phasellus", "tincidunt",
			"ante", "ac", "malesuada", "fringilla", "Aliquam", "ac", "erat", "tempus", "lacinia", "ex", "quis", "sagittis",
			"ipsum", "Nam", "neque", "augue", "rhoncus", "quis", "volutpat", "rhoncus", "sodales", "ut", "dolor",
			"Quisque", "at", "massa", "at", "diam", "suscipit", "posuere", "quis", "eget", "neque", "Phasellus",
			"eu", "euismod", "lacus", "Vivamus", "ac", "magna", "feugiat", "ullamcorper", "ex", "eget", "aliquam",
			"nisi", "Suspendisse", "aliquet", "nunc", "in", "elit", "facilisis", "ultrices", "Suspendisse", "non",
			"vehicula", "urna", "pulvinar", "fermentum", "diam", "Vivamus", "convallis", "leo", "sit", "amet",
			"turpis", "facilisis", "convallis", "Integer", "finibus", "enim", "odio", "Curabitur", "euismod",
			"nisi", "nulla", "vestibulum", "cursus", "diam", "ornare", "ut", "Nunc", "feugiat", "efficitur",
			"vehicula", "Class", "aptent", "taciti", "sociosqu", "ad", "litora", "torquent", "per", "conubia",
			"nostra", "per", "inceptos", "himenaeos", "Sed", "mattis", "ligula", "non", "velit", "ultricies",
			"lobortis", "Integer", "finibus", "tortor", "et", "velit", "accumsan", "congue", "Nam", "consectetur",
			"lacus", "efficitur", "sodales", "sollicitudin", "Phasellus", "non", "tempor", "velit", "Nullam", "id",
			"accumsan", "lacus", "Quisque", "est", "felis", "placerat", "in", "nunc", "et", "tempor", "scelerisque",
			"enim", "Quisque", "malesuada", "sodales", "lectus", "sit", "amet", "facilisis", "massa", "tincidunt",
			"a", "Suspendisse", "potenti", "Etiam", "vitae", "metus", "vehicula", "sagittis", "lacus", "a", "viverra",
			"purus", "Donec", "convallis", "eros", "mollis", "rhoncus", "magna", "id", "vulputate", "magna", "Donec",
			"consectetur", "in", "erat", "vitae", "sollicitudin", "Vestibulum", "facilisis", "et", "ligula", "sit",
			"amet", "tristique", "Ut", "dapibus", "finibus", "arcu", "ac", "rhoncus", "arcu", "placerat", "in",
			"Sed", "a", "venenatis", "augue", "Quisque", "quis", "sapien", "vel", "quam", "congue", "feugiat",
			"non", "eget", "mauris", "Vivamus", "ac", "elit", "eu", "nunc", "blandit", "tempus", "Nulla", "orci",
			"est", "blandit", "vehicula", "ipsum", "vitae", "molestie", "venenatis", "ante", "Ut", "ornare", "tempor",
			"est", "eget", "aliquet", "magna", "blandit", "vitae", "Phasellus", "in", "erat", "augue", "Praesent",
			"lacinia", "augue", "nibh", "ac", "feugiat", "nisi", "blandit", "nec", "Quisque", "nulla", "lacus",
			"molestie", "ut", "dignissim", "a", "pulvinar", "non", "sem", "Curabitur", "ut", "turpis", "pretium",
			"fringilla", "risus", "vitae", "tempor", "felis", "Phasellus", "at", "pharetra", "nisl", "Curabitur",
			"lobortis", "enim", "dui", "et", "cursus", "orci", "porta", "ut", "Nullam", "et", "vulputate", "neque",
			"Phasellus", "non", "convallis", "lectus", "Pellentesque", "semper", "cursus", "euismod", "Suspendisse",
			"aliquam", "aliquam", "magna", "ut", "euismod", "leo", "tincidunt", "ut"}),
		CreateSetOfStrings([]string{"Lorem", "adipiscing", "Aenean", "faucibus", "Morbi", "Mauris", "iaculis",
			"libero", "porttitor", "dictum", "sed", "Duis", "laoreet", "egestas", "gravida", "luctus", "hendrerit",
			"Interdum", "fames", "primis", "facilisi", "mi"}),
	},
}

func setTestBlock(t *testing.T) {
	for _, tCase := range testCases {
		intersect := Intersection(tCase.set1, tCase.set2)
		if !EqualSets(intersect, tCase.interRes) {
			t.Error(tCase.set1, "ᴖ", tCase.set2, "==", tCase.interRes, "!=", intersect)
		}
		union := Union(tCase.set1, tCase.set2)
		if !EqualSets(union, tCase.unionRes) {
			t.Error(tCase.set1, "ᴗ", tCase.set2, "==", tCase.unionRes, "!=", union)
		}
		diff := Difference(tCase.set1, tCase.set2)
		if !EqualSets(diff, tCase.difRes) {
			t.Error(tCase.set1, "\\", tCase.set2, "==", tCase.difRes, "!=", diff)
		}
	}
}

func TestSetOperations(t *testing.T) {
	for i := 0; i < 50000; i++ {
		setTestBlock(t)
	}
}
