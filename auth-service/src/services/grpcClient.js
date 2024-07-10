import grpc from '@grpc/grpc-js';
import { loadSync } from '@grpc/proto-loader';
import { join } from 'path';


const PROTO_PATH = join(__dirname, "../../../proto/service.proto");

const packageDefinition = loadSync(PROTO_PATH, {

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