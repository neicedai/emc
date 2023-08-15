for ((i=$1;i<=$2;i++)) do
cd /home
cp -r emc emc$i
cd emc$i
ehco"init EMC$i"
./edge-matrix secrets init --data-dir edge_data
sleep 5s
echo "Start emc"
./start.sh
sleep 5s
./edge-matrix miner status
