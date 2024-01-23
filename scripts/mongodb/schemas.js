const db = connect(process.env.MONGODB_URL || 'mongodb://localhost:27017/carriers');

const collection = db.getCollection('carriers');
if (collection) {
    print("Collection already created");
    quit();
}
db.createCollection('carriers', {
    validator: {
        $jsonSchema: {
            bsonType: 'object',
            required: ['name', 'dba'],
            properties: {
                name: {
                    bsonType: 'string',
                    description: 'must be a string and is required'
                },
                dba: {
                    bsonType: 'string',
                    description: 'must be a string and is required'
                }
            }
        }
    }
});