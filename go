for ((i=$1;i<=$2;i++)) do
cd /home
cp -r emc emc$i
cd emc$i
ehco"init EMC$i"
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
sed -i 's#7860#$p5'#g' emc$i
cp emc$i /usr/local/bin

