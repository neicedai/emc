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
p1=`expr $i + 41000`
p2=`expr $i + 42000`
p3=`expr $i + 43000`
p4=`expr $i + 44000`
p5=`expr $i + 45000`
sed -i 's#0.0.0.0:50000#0.0.0.0:'$p1'#g' emc$i
sed -i 's#0.0.0.0:50001#0.0.0.0:'$p2'#g' emc$i
sed -i 's#0.0.0.0:50002#0.0.0.0:'$p3'#g' emc$i
sed -i 's#0.0.0.0:50003#0.0.0.0:'$p4'#g' emc$i
sed -i 's#7860#'$p5'#g' emc$i
sed -i 's#emcth#emc'$i'#g' emc$i
chmod 777 emc$i
cp emc$i /usr/local/bin
service emc$i start
echo "waiting for register"
sleep 20s
cd /home/emc$i
./edge-matrix miner register --commit set --node computing --principal rakqt-iyuu2-4hre4-y4its-xk7q4-stmin-ptusv-yqagc-gnofs-ct2ev-fqe --grpc-address "127.0.0.1:$p1"
sleep 60s
done
