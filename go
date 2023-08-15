for ((i=$1;i<=$2;i++)) do
cd /home
cp -r emc emc$i
cd emc$i
ehco"init EMC$i"
./edge-matrix secrets init --data-dir edge_data
sleep 5s
echo "Start emc"
sed -i 's#btfsl#emc'$i'#g' /etc/systemd/system/btfs$i.service 
./start.sh
sleep 5s
./edge-matrix miner status
