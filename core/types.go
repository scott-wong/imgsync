package core

import "fmt"

type Image struct {
	Repo string
	User string
	Name string
	Tag  string
}

func (img Image) String() string {
	if img.User != "" {
		return fmt.Sprintf("%s/%s/%s:%s", img.Repo, img.User, img.Name, img.Tag)
	}
	return fmt.Sprintf("%s/%s:%s", img.Repo, img.Name, img.Tag)
}

func (img Image) MergeName() string {
	if img.User != "" {
		return fmt.Sprintf("%s_%s_%s", img.Repo, img.User, img.Name)
	}
	return fmt.Sprintf("%s_%s", img.Repo, img.Name)
}

type Images []Image

func (imgs Images) Len() int           { return len(imgs) }
func (imgs Images) Less(i, j int) bool { return imgs[i].String() < imgs[j].String() }
func (imgs Images) Swap(i, j int)      { imgs[i], imgs[j] = imgs[j], imgs[i] }

type Manifest struct {
	Name      string
	JSONValue string
}

type Manifests []Manifest

func (m Manifests) Len() int           { return len(m) }
func (m Manifests) Less(i, j int) bool { return m[i].Name < m[j].Name }
func (m Manifests) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }