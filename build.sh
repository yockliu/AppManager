#! /bin/sh

echo "run build sh"

projectPath=$1
gitTag=$2
channelGroup=$3

echo projectPath + $projectPath
echo gitTag + $gitTag
echo channelGroup + $channelGroup

cd $projectPath

git reset --hard
git pull
git co $gitTag

#sed -i '' '/productFlavors/a\ 
#aaa
#' zhoumo/build.gradle

#sed -i '' 's/productFlavors {*}/productFlavors{\n}/g' zhoumo/build.gradle

#sed -i '' '/productFlavors \{/{N;/!\}/}D' zhoumo/build.gradle

rm  tmp.gradle

goinFlavors=false
DONE=false
until $DONE 
do read || DONE=true
	if [ $goinFlavors = false ]; then
		echo "$REPLY" >> tmp.gradle
	fi
	if [ "$REPLY" = "    productFlavors {" ]; then
		goinFlavors=true
	fi
	if [ $goinFlavors = true ] && [ "$REPLY" = "    }" ]; then
		for channel in `echo "\t\t$channelGroup" | tr ',' ' '`
		do
			#echo $channel >> tmp.gradle
			echo $channel | sed 's/and-//' >> tmp.gradle
		done
		echo $REPLY >> tmp.gradle
		goinFlavors=false
	fi
done < zhoumo/build.gradle

mv tmp.gradle zhoumo/build.gradle

./gradlew assembleRelease 
