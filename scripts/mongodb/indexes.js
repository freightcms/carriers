const db = connect(process.env.MONGODB_URL || 'mongodb://localhost:27017/carriers');

const indexes = db.carriers.getIndexes();
if (indexes.length > 1) {
    print("Indexes already created");
    quit();
}
// Create indexes
if (!indexes.find(i => i.name === "ux_carriers_name")) {
    db.carriers.createIndex({ name: 1 }, { 
        name: "ux_carriers_name",
        unique: true 
    });
}
if (!indexes.find(i => i.name === "ux_carriers_dba")) {
    db.carriers.createIndex({ dba: 1 }, {
        name: "ux_carriers_dba",
        unique: true, 
    });
}