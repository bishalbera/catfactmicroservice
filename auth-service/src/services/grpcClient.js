const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const path = require('path');



const PROTO_PATH = path.join(__dirname, "../../../proto/service.proto");

const packageDefinition = protoLoader.loadSync(PROTO_PATH, {

    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
});

const catfactProto = grpc.loadPackageDefinition(packageDefinition);

const client = new catfactProto.CatFact("localhost:50051",grpc.credentials.createInsecure());

const getCatFact = (callback) => {

    client.GetCatFact({}, (error, response) => {
        if (error) {
            callback(error, null);
        } else{
            callback(null, response);
        }
    });
};

module.exports = { getCatFact };