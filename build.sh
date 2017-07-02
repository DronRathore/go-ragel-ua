echo "Checking golang"
go_cmd=$(which go)
if [[ $? != 0 ]]
  then
    echo "Golang binary not found, trying to install..."
    if [[ "${platform,,}" == "darwin" ]]
    then
      out=$(brew install golang)
    else
      out=$(yum install golang)
  fi
  if [[ $? != 0 ]]
    then
    echo "Failed to install golang, please check the above message"
    echo "${out}"
    exit $?
  fi
else
  echo "Using golang from ${go_cmd}"
fi

# Check ragel's presense
echo "Checking ragel..."
ragel_cmd=$(which ragel)
rc=$?
platform=$(uname)
helper="src/ua"

if [[ $rc != 0 ]]; then
  echo "Installing Ragel..."
  if [[ "${platform,,}" == "darwin" ]]
    then
      out=$(brew install ragel)
    else
      out=$(yum install ragel)
  fi
  if [[ $? != 0 ]]
    then
    echo "Failed to install ragel, please check the above message"
    echo "${out}"
    exit $?
  fi
else
  echo "Ragel found in ${ragel_cmd} \n"
fi

set -e

echo "Building .rl files in ${helper} folder"
if [ "$GOENV" = "production" ]; then
  echo -e "Optimise code generation set to true in ${GOENV}\n"
  ragel_opts='-Z -G2'
else
  ragel_opts='-Z'
  echo "Unoptimise ragel code generation in dev mode"
fi
ragel ${ragel_opts} ./is_mobile.rl -o ${helper}/is_mobile.go
ragel ${ragel_opts} ./ua_version.rl -o ${helper}/ua_version.go
ragel ${ragel_opts} ./is_bot.rl -o ${helper}/is_bot.go
echo "Ragel build successful\n"
echo "==================================\n"
