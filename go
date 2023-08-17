for ((i=$1;i<=$2;i++)) do
cd /home
cp -r emc emc$i
cd emc$i
echo "init EMC$i"
./edge-matrix secrets init --data-dir edge_data
sleep 5s
echo "Start emc"
cd /root/emc
cp emc.service /etc/systemd/system/emc$i.service
sed -i 's#emcth#emc'$i'#g' /etc/systemd/system/emc$i.service
cp emc emc$i
sed -i 's#emcth#emc'$i'#g' emc$i
chmod 777 emc$i
cp emc$i /usr/local/bin
service emc$i start
echo "waiting for register"
sleep 20s
cd /home/emc$i
./edge-matrix miner register --commit set --node computing --principal rakqt-iyuu2-4hre4-y4its-xk7q4-stmin-ptusv-yqagc-gnofs-ct2ev-fqe 
sleep 60s
done
