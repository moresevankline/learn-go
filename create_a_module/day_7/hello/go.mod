module example.com/hello

go 1.24.4

replace example.com/greetings => ../greetings

require example.com/greetings v1.1.0

// To reference a published module, a go.mod file would typically omit the replace directive and use a require directive with a tagged version number at the end.

// require example.com/greetings v0.0.0-00010101000000-000000000000 -> require example.com/greetings v1.1.0
