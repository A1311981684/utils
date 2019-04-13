package tar

import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func UnTar(tarFileName string, unTarFilePath string) error  {

	tarFile, err := os.Open(tarFileName)
	if err != nil {
		return err
	}
	defer func() {
		err := tarFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	tarReader := tar.NewReader(tarFile)

	path, err := os.Stat(unTarFilePath)
	if path != nil{
		//If directory already exists, remove it (include everything under)
		err = os.RemoveAll(unTarFilePath)
		if err != nil {
			return err
		}
	}
	err = os.Mkdir(unTarFilePath, os.ModePerm)
	if err != nil {
		return err
	}
	//Un-tar every single file
	for{
		hdr, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fileName := unTarFilePath + hdr.Name
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, tarReader)
		if err != nil{
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
