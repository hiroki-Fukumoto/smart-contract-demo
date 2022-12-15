FROM node:alpine

RUN npm install -g ganache-cli
RUN mkdir -p /var/ganache

ENV DOCKER true
ENV DATADIR /var/ganache
ENV PRIVATEKEY 0x928f6acb4f4305c44d86b117d4a72d559cddeacdc801b6cde6a9c7f1a63f0215
ENV STARTINGWEI 10000000000000000000000

ENV MNEMONIC pillows andymori al tomoyuki spitz abc quruli whoops fukurouz air bluehearts highlows
ENV NETWORKID 5777

EXPOSE 8545

# ENTRYPOINT ["sh", "-c", "ganache-cli --host=0.0.0.0 --db=${DATADIR} --mnemonic=\"${MNEMONIC}\" --networkId=${NETWORKID}"]
ENTRYPOINT ["sh", "-c", "ganache-cli --host=0.0.0.0 --account=${PRIVATEKEY},${STARTINGWEI} --mnemonic=\"${MNEMONIC}\" --networkId=${NETWORKID}"]
