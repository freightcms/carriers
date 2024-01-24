const db = connect(process.env.MONGODB_URL || 'mongodb://localhost:27017/carriers');

const collection = db.getCollection('carriers');
if (collection) {
    collection.drop();
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

const identifyingCodeCollection = db.getCollection('identifying_codes');
if (identifyingCodeCollection) {
    identifyingCodeCollection.drop();
}
console.log("Creating identifying_codes collection");
db.createCollection('identifying_codes', {
    validator: {
        $jsonSchema: {
            bsonType: 'object',
            required: ['code', 'type', 'carrier_id'],
            properties: {
                code: {
                    bsonType: 'string',
                    description: 'must be a string and is required'
                },
                type: {
                    bsonType: 'string',
                    description: 'must be a string and is required'
                },
                carrier_id: {
                    bsonType: 'objectId',
                    description: 'must be a objectId and is required'
                }
            }
        }
    }
});


const contactCollection = db.getCollection('contacts');
if (contactCollection) {
    contactCollection.drop();
}
console.log("Creating contacts collection");
db.createCollection('contacts', {
    validator: {
        $jsonSchema: {
            bsonType: 'object',
            required: ['name', 'phone', 'email', 'carrier_id'],
            properties: {
                name: {
                    bsonType: 'string',
                    description: 'must be a string and is required'
                },
                phone: {
                    bsonType: 'string',
                    description: 'must be a string and is required'
                },
                email: {
                    bsonType: 'string',
                    description: 'must be a string and is required'
                },
                fax: {
                    bsonType: 'string',
                    description: 'must be a string'
                },
                carrier_id: {
                    bsonType: 'objectId',
                    description: 'must be a objectId and is required'
                }
            }
        }
    }
});

const addressCollection = db.getCollection("addresses");
if (addressCollection) {
    addressCollection.drop();
}
console.log("Creating addresses collection");
db.createCollection('addresses', {
    validator: {
        $jsonSchema: {
            bsonType: 'object',
            required: ['type', 'line_1', 'region', 'local', 'zip', 'country', 'carrier_id'],
            properties: {
                type: {
                    bsonType: 'string',
                    enum: ['billing', 'physical', 'mailing'],
                    description: 'must be a string and is required'
                },
                line_1: {
                    bsonType: 'string',
                    description: 'line_1 is typically used for street address and is required'
                },
                line_2: {
                    bsonType: 'string',
                    description: 'line_2 is typically used for suite numbers and is optional'
                },
                line_3: {
                    bsonType: 'string',
                    description: 'line_3 is typically used for floor numbers and is optional'
                },
                region: {
                    bsonType: 'string',
                    description: 'region is the state and is required'
                },
                local: {
                    bsonType: 'string',
                    description: 'local is the city and is required'
                },
                zip: {
                    bsonType: 'string',
                    description: 'zip and/or postal code must be a string and is required'
                },
                country: {
                    bsonType: 'string',
                    description: 'country must be a string and is required. 2 letter ISO 3166-1 alpha-2 country code.'
                },
                carrier_id: {
                    bsonType: 'objectId',
                    description: 'must be a objectId and is required'
                }
            }
        }
    }
});
