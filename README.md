# raw2wav
RAW 음성 파일을 PCM/G711 WAV로 변환하는 유틸입니다. VCE 파일을 변환하는데 유용합니다.  
golang으로 작성 되었습니다. https://golang.org/ 에서 다운로드 받아 go를 설치 해주세요.  
ffmpeg을 이용한 파일 변환입니다. https://ffmpeg.org/download.html 에서 플랫폼에 맞는 ffmpeg 바이너리를 다운로드 받아 ffpmeg을 설치 해주세요. 
Linux/Windows 지원됩니다.   

빌드방법  
$> go build raw2wav.go  

준비사항  
빌드된 실행 파일과 ffmpeg 실행파일을 동일한 경로에 위치하게 해 주세요.  

사용방법  
Usage: raw2wav [input type] [dir path] [output_type]  
    - input type: alaw|mulaw|s16le|u16le  
    - output type: pcm_alaw|pcm_mulaw|pcm_s16le|pcm_u16le  
        ex)raw2wav alaw ./sound/*.vce pcm_alaw  
