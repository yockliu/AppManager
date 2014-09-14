#! /bin/sh

echo "run build sh"

SH_PATH=`pwd`
APP_IDENTI=$1
projectPath=$2
gitTag=$3
channelGroup=$4

tmpgradle=tmp.gradle
buildgradle=zhoumo/build.gradle

echo projectPath + $projectPath
echo gitTag + $gitTag
echo channelGroup + $channelGroup

cd $projectPath

git reset --hard
git pull
git checkout $gitTag

rm $tmpgradle

goinFlavors=false
DONE=false
until $DONE 
do read || DONE=true
	if [ $goinFlavors = false ]; then
		echo "$REPLY" >> $tmpgradle
	fi
	if [ "$REPLY" = "    productFlavors {" ]; then
		goinFlavors=true
	fi
	if [ $goinFlavors = true ] && [ "$REPLY" = "    }" ]; then
		for channel in `echo "$channelGroup" | tr ',' ' '`
		do
			#echo $channel >> tmp.gradle
			echo $channel | sed 's/and-//' >> $tmpgradle
		done
		echo $REPLY >> $tmpgradle
		goinFlavors=false
	fi
done < $buildgradle

mv $tmpgradle $buildgradle

./gradlew assembleRelease 

outputSource=$projectPath/zhoumo/build/outputs/apk/
outputDest=$SH_PATH/static/apk/$APP_IDENTI/$gitTag/
echo $outputSource
echo $outputDest
mkdir -p $outputDest
rsync -vaz --exclude="*unaligned.apk" $outputSource  $outputDest
